package repository

import (
	"fmt"
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreateLeague(league *models.League) error {
	return database.DB.Create(league).Error
}

// GetLeagues retorna todas as ligas associadas a um usuário específico
func GetLeagues(userID string) ([]models.League, error) {
	var leagues []models.League
	err := database.DB.Where("user_id = ?", userID).Find(&leagues).Error
	return leagues, err
}

// GetLeagueByID retorna uma liga específica pelo ID, desde que pertença ao usuário
func GetLeagueByID(userID, leagueID string) (models.League, error) {
	var league models.League
	err := database.DB.Where("id = ? AND user_id = ?", leagueID, userID).First(&league).Error
	return league, err
}

// UpdateLeague atualiza uma liga específica, desde que pertença ao usuário
func UpdateLeague(userID, leagueID string, updates map[string]interface{}) (models.League, error) {
	var league models.League
	if err := database.DB.Model(&league).Where("id = ? AND user_id = ?", leagueID, userID).Updates(updates).Error; err != nil {
		return models.League{}, err
	}
	database.DB.First(&league, "id = ? AND user_id = ?", leagueID, userID)
	return league, nil
}

// DeleteLeague deleta uma liga específica, desde que pertença ao usuário
func DeleteLeague(userID, leagueID string) error {
	return database.DB.Delete(&models.League{}, "id = ? AND user_id = ?", leagueID, userID).Error
}

// GetTeamsAndGamesByLeagueID retorna times e jogos de uma liga específica, desde que pertença ao usuário
func GetTeamsAndGamesByLeagueID(userID, leagueID string) ([]models.Team, []models.Game, error) {
	var teams []models.Team
	var games []models.Game

	// Verifica se a liga pertence ao usuário
	var league models.League
	if err := database.DB.Where("id = ? AND user_id = ?", leagueID, userID).First(&league).Error; err != nil {
		return nil, nil, err
	}

	// Buscar times da liga
	if err := database.DB.Where("league_id = ?", leagueID).Find(&teams).Error; err != nil {
		return nil, nil, err
	}

	// Buscar jogos concluídos da liga
	if err := database.DB.Where("league_id = ? AND status = ?", leagueID, "completed").Find(&games).Error; err != nil {
		return nil, nil, err
	}

	return teams, games, nil
}

type PlayerStatistics struct {
	PlayerID    string `json:"player_id"`
	PlayerName  string `json:"player_name"`
	TotalPoints int    `json:"total_points"`
}

// GetTotalPointsRanking retorna o ranking de pontos de uma liga específica, desde que pertença ao usuário
func GetTotalPointsRanking(userID, leagueID string, limit int) ([]PlayerStatistics, error) {
	// Verifica se a liga pertence ao usuário
	var league models.League
	if err := database.DB.Where("id = ? AND user_id = ?", leagueID, userID).First(&league).Error; err != nil {
		return nil, err
	}

	// Query SQL para buscar o ranking
	const query = `
		SELECT 
			gs.player_id,
			p.name AS player_name,
			SUM(gs.points) AS total_points
		FROM game_statistics gs
		JOIN players p ON gs.player_id = p.id
		JOIN teams t ON p.team_id = t.id
		WHERE t.league_id = ?
		GROUP BY gs.player_id, p.name
		ORDER BY total_points DESC
		LIMIT ?`

	// Executar a query
	var topScorers []PlayerStatistics
	if err := database.DB.Raw(query, leagueID, limit).Scan(&topScorers).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar ranking de pontos: %w", err)
	}

	return topScorers, nil
}

type PlayerThreePointsStatistics struct {
	PlayerID    string `json:"player_id"`
	PlayerName  string `json:"player_name"`
	ThreePoints int    `json:"three_points"`
}

// GetTotalThreePointsRanking retorna o ranking de pontos de três de uma liga específica, desde que pertença ao usuário
func GetTotalThreePointsRanking(userID, leagueID string, limit int) ([]PlayerThreePointsStatistics, error) {
	// Verifica se a liga pertence ao usuário
	var league models.League
	if err := database.DB.Where("id = ? AND user_id = ?", leagueID, userID).First(&league).Error; err != nil {
		return nil, err
	}

	// Query SQL para buscar o ranking
	const query = `
		SELECT 
			gs.player_id,
			p.name AS player_name,
			SUM(gs.three_points) AS three_points
		FROM game_statistics gs
		JOIN players p ON gs.player_id = p.id
		JOIN teams t ON p.team_id = t.id
		WHERE t.league_id = ?
		GROUP BY gs.player_id, p.name
		ORDER BY three_points DESC
		LIMIT ?`

	// Executar a query
	var topScorers []PlayerThreePointsStatistics
	if err := database.DB.Raw(query, leagueID, limit).Scan(&topScorers).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar ranking de pontos: %w", err)
	}

	return topScorers, nil
}
