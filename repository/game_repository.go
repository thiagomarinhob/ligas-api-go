package repository

import (
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreateGame(game *models.Game) error {
	return database.DB.Create(game).Error
}

func GetGames() ([]models.Game, error) {
	var games []models.Game
	err := database.DB.
		Preload("League").
		Preload("TeamA").
		Preload("TeamB").
		Find(&games).Error
	return games, err
}

func GetGameByID(id string) (models.Game, error) {
	var game models.Game
	err := database.DB.
		Preload("League").
		Preload("TeamA").
		Preload("TeamB").
		First(&game, "id = ?", id).Error
	return game, err
}

func UpdateGame(id string, updates map[string]interface{}) (models.Game, error) {
	var game models.Game
	if err := database.DB.Model(&game).Where("id = ?", id).Updates(updates).Error; err != nil {
		return models.Game{}, err
	}
	database.DB.First(&game, "id = ?", id)
	return game, nil
}

func DeleteGame(id string) error {
	return database.DB.Delete(&models.Game{}, "id = ?", id).Error
}
