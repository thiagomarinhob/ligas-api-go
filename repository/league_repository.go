package repository

import (
	"fmt"
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreateLeague(league *models.League) error {
	return database.DB.Create(league).Error
}

func GetLeagues() ([]models.League, error) {
	var leagues []models.League
	err := database.DB.Find(&leagues).Error
	return leagues, err
}

func GetLeagueByID(id string) (models.League, error) {
	var league models.League
	err := database.DB.First(&league, "id = ?", id).Error
	return league, err
}

func UpdateLeague(id string, updates map[string]interface{}) (models.League, error) {
	var league models.League
	if err := database.DB.Model(&league).Where("id = ?", id).Updates(updates).Error; err != nil {
		return models.League{}, err
	}
	database.DB.First(&league, "id = ?", id)
	return league, nil
}

func DeleteLeague(id string) error {
	return database.DB.Delete(&models.League{}, "id = ?", id).Error
}

func GetTeamsAndGamesByLeagueID(leagueID string) ([]models.Team, []models.Game, error) {
	var teams []models.Team
	var games []models.Game

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

func GetTotalPointsRanking(leagueID string, limit int) ([]PlayerStatistics, error) {

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

func GetTotalThreePointsRanking(leagueID string, limit int) ([]PlayerThreePointsStatistics, error) {

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
