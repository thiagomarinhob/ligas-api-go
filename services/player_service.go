package services

import (
	"errors"
	"go-api-ligas/models"
	"go-api-ligas/repository"
	"time"
)

func CreatePlayer(input struct {
	Name      string  `json:"name"`
	Number    int     `json:"number"`
	TeamID    string  `json:"team_id"`
	Position  string  `json:"position"`
	Height    float32 `json:"height"`
	Weight    float32 `json:"weight"`
	BirthDate *string `json:"birth_date"`
	Photo     string  `json:"photo"`
	Active    bool    `json:"active"`
}) (models.Player, error) {
	var birthDate *time.Time
	if input.BirthDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.BirthDate)
		if err != nil {
			return models.Player{}, errors.New("invalid birth date format")
		}
		birthDate = &parsedDate
	}

	player := models.Player{
		Name:      input.Name,
		Number:    input.Number,
		TeamID:    input.TeamID,
		Position:  input.Position,
		Height:    input.Height,
		Weight:    input.Weight,
		BirthDate: birthDate,
		Photo:     input.Photo,
		Active:    input.Active,
	}

	err := repository.CreatePlayer(&player)
	return player, err
}

func GetPlayers() ([]models.Player, error) {
	return repository.GetPlayers()
}

func GetPlayerByID(id string) (models.Player, error) {
	return repository.GetPlayerByID(id)
}

func UpdatePlayer(id string, updates map[string]interface{}) (models.Player, error) {
	return repository.UpdatePlayer(id, updates)
}

func DeletePlayer(id string) error {
	return repository.DeletePlayer(id)
}
