package repository

import (
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreatePlayer(player *models.Player) error {
	return database.DB.Create(player).Error
}

func GetPlayers() ([]models.Player, error) {
	var players []models.Player
	err := database.DB.Preload("Team").Find(&players).Error
	return players, err
}

func GetPlayerByID(id string) (models.Player, error) {
	var player models.Player
	err := database.DB.Preload("Team").First(&player, "id = ?", id).Error
	return player, err
}

func UpdatePlayer(id string, updates map[string]interface{}) (models.Player, error) {
	var player models.Player
	if err := database.DB.Model(&player).Where("id = ?", id).Updates(updates).Error; err != nil {
		return models.Player{}, err
	}
	database.DB.First(&player, "id = ?", id)
	return player, nil
}

func DeletePlayer(id string) error {
	return database.DB.Delete(&models.Player{}, "id = ?", id).Error
}
