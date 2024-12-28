package services

import (
	"errors"
	"go-api-ligas/models"
	"go-api-ligas/repository"
	"time"
)

func CreateTeam(input struct {
	Name           string  `json:"name"`
	Acronym        string  `json:"acronym"`
	LeagueID       string  `json:"league_id"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	FoundationDate *string `json:"foundation_date"`
	Badge          string  `json:"badge"`
	Description    string  `json:"description"`
}) (models.Team, error) {
	var foundationDate *time.Time
	if input.FoundationDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.FoundationDate)
		if err != nil {
			return models.Team{}, errors.New("invalid foundation date format")
		}
		foundationDate = &parsedDate
	}

	team := models.Team{
		Name:           input.Name,
		Acronym:        input.Acronym,
		LeagueID:       input.LeagueID,
		City:           input.City,
		State:          input.State,
		FoundationDate: foundationDate,
		Badge:          input.Badge,
		Description:    input.Description,
	}

	err := repository.CreateTeam(&team)
	return team, err
}

func GetTeams() ([]models.Team, error) {
	return repository.GetTeams()
}

func GetTeamByID(id string) (models.Team, error) {
	return repository.GetTeamByID(id)
}

func UpdateTeam(id string, updates map[string]interface{}) (models.Team, error) {
	return repository.UpdateTeam(id, updates)
}

func DeleteTeam(id string) error {
	return repository.DeleteTeam(id)
}
