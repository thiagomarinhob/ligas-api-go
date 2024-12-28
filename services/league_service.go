package services

import (
	"errors"
	"go-api-ligas/models"
	"go-api-ligas/repository"
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
