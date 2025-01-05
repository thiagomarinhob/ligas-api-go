package services

import (
	"errors"
	"go-api-ligas/models"
	"go-api-ligas/repository"
	"sort"
	"time"
)

func CreateLeague(input struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Gender      string  `json:"gender"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	MaxPlayers  int     `json:"max_players"`
	MaxTeams    int     `json:"max_teams"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	UserID      string  `json:"user_id"`
}) (models.League, error) {
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return models.League{}, errors.New("invalid start date format")
	}

	var endDate *time.Time
	if input.EndDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.EndDate)
		if err != nil {
			return models.League{}, errors.New("invalid end date format")
		}
		endDate = &parsedDate
	}

	league := models.League{
		Name:        input.Name,
		Category:    input.Category,
		Gender:      input.Gender,
		StartDate:   startDate,
		EndDate:     endDate,
		MaxPlayers:  input.MaxPlayers,
		MaxTeams:    input.MaxTeams,
		Status:      input.Status,
		Description: input.Description,
		UserID:      input.UserID,
	}

	err = repository.CreateLeague(&league)
	return league, err
}

func GetLeagues() ([]models.League, error) {
	return repository.GetLeagues()
}

func GetLeagueByID(id string) (models.League, error) {
	return repository.GetLeagueByID(id)
}

func UpdateLeague(id string, updates map[string]interface{}) (models.League, error) {
	return repository.UpdateLeague(id, updates)
}

func DeleteLeague(id string) error {
	return repository.DeleteLeague(id)
}

func GetLeagueStandings(leagueID string) ([]map[string]interface{}, error) {
	// Buscar times e jogos associados à liga
	teams, games, err := repository.GetTeamsAndGamesByLeagueID(leagueID)
	if err != nil {
		return nil, err
	}

	// Mapa para armazenar vitórias e derrotas
	standings := make(map[string]map[string]interface{})
	for _, team := range teams {
		standings[team.ID] = map[string]interface{}{
			"Team": map[string]interface{}{
				"id":      team.ID,
				"name":    team.Name,
				"acronym": team.Acronym,
				"badge":   team.Badge,
			},
			"Points":        0,
			"Wins":          0,
			"Losses":        0,
			"PointsScored":  0,
			"PointsAgainst": 0,
			"Balance":       0,
		}
	}

	// Calcular vitórias e derrotas com base nos pontos dos jogos
	for _, game := range games {
		teamAStanding := standings[game.TeamAID]
		teamBStanding := standings[game.TeamBID]

		// Atualizar pontos marcados e sofridos
		teamAStanding["PointsScored"] = teamAStanding["PointsScored"].(int) + game.PointsTeamA
		teamAStanding["PointsAgainst"] = teamAStanding["PointsAgainst"].(int) + game.PointsTeamB
		teamBStanding["PointsScored"] = teamBStanding["PointsScored"].(int) + game.PointsTeamB
		teamBStanding["PointsAgainst"] = teamBStanding["PointsAgainst"].(int) + game.PointsTeamA

		// Atualizar saldo de pontos
		teamAStanding["Balance"] = teamAStanding["PointsScored"].(int) - teamAStanding["PointsAgainst"].(int)
		teamBStanding["Balance"] = teamBStanding["PointsScored"].(int) - teamBStanding["PointsAgainst"].(int)

		// Determinar resultados e atualizar standings
		if game.PointsTeamA > game.PointsTeamB {
			teamAStanding["Wins"] = teamAStanding["Wins"].(int) + 1
			teamAStanding["Points"] = teamAStanding["Points"].(int) + 3 // Pontos por vitória
			teamBStanding["Losses"] = teamBStanding["Losses"].(int) + 1
		} else if game.PointsTeamA < game.PointsTeamB {
			teamBStanding["Wins"] = teamBStanding["Wins"].(int) + 1
			teamBStanding["Points"] = teamBStanding["Points"].(int) + 3
			teamAStanding["Losses"] = teamAStanding["Losses"].(int) + 1
		}
	}

	// Converter o mapa para uma lista
	result := []map[string]interface{}{}
	for _, stats := range standings {
		result = append(result, stats)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i]["Points"].(int) == result[j]["Points"].(int) {
			return result[i]["Balance"].(int) > result[j]["Balance"].(int)
		}
		return result[i]["Points"].(int) > result[j]["Points"].(int)
	})

	return result, nil
}

func GetTotalPointsRanking(leagueID string, limit int) ([]repository.PlayerStatistics, error) {
	if leagueID == "" {
		return nil, errors.New("league ID is required")
	}

	if limit == 0 {
		limit = 10
	}

	return repository.GetTotalPointsRanking(leagueID, limit)
}

func GetTotalThreePointsRanking(leagueID string, limit int) ([]repository.PlayerThreePointsStatistics, error) {
	if leagueID == "" {
		return nil, errors.New("league ID is required")
	}

	if limit == 0 {
		limit = 10
	}

	return repository.GetTotalThreePointsRanking(leagueID, limit)
}
