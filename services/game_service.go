package services

import (
	"errors"
	"go-api-ligas/models"
	"go-api-ligas/repository"
	"time"
)

func CreateGame(input struct {
	LeagueID    string `json:"league_id"`
	TeamAID     string `json:"team_a_id"`
	TeamBID     string `json:"team_b_id"`
	DateTime    string `json:"date_time"`
	Location    string `json:"location"`
	PointsTeamA int    `json:"points_team_a"`
	PointsTeamB int    `json:"points_team_b"`
	Status      string `json:"status"`
	Description string `json:"description"`
}) (models.Game, error) {
	dateTime, err := time.Parse("2006-01-02T15:04:05", input.DateTime)
	if err != nil {
		return models.Game{}, errors.New("invalid date format")
	}

	game := models.Game{
		LeagueID:    input.LeagueID,
		TeamAID:     input.TeamAID,
		TeamBID:     input.TeamBID,
		DateTime:    dateTime,
		Location:    input.Location,
		PointsTeamA: input.PointsTeamA,
		PointsTeamB: input.PointsTeamB,
		Status:      input.Status,
		Description: input.Description,
	}

	err = repository.CreateGame(&game)
	return game, err
}

func GetGames() ([]models.Game, error) {
	return repository.GetGames()
}

func GetGameByID(id string) (models.Game, error) {
	return repository.GetGameByID(id)
}

func UpdateGame(id string, updates map[string]interface{}) (models.Game, error) {
	return repository.UpdateGame(id, updates)
}

func DeleteGame(id string) error {
	return repository.DeleteGame(id)
}
