package repository

import (
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreateGameStatistics(statistics *models.GameStatistics) error {
	return database.DB.Create(statistics).Error
}

func GetGameStatistics() ([]models.GameStatistics, error) {
	var statistics []models.GameStatistics
	err := database.DB.Preload("Game").Preload("Player").Find(&statistics).Error
	return statistics, err
}

func GetGameStatisticsByID(id string) (models.GameStatistics, error) {
	var statistics models.GameStatistics
	err := database.DB.Preload("Game").Preload("Player").First(&statistics, "id = ?", id).Error
	return statistics, err
}

func UpdateGameStatistics(id string, updates map[string]interface{}) (models.GameStatistics, error) {
	var statistics models.GameStatistics
	if err := database.DB.Model(&statistics).Where("id = ?", id).Updates(updates).Error; err != nil {
		return models.GameStatistics{}, err
	}
	database.DB.First(&statistics, "id = ?", id)
	return statistics, nil
}

func DeleteGameStatistics(id string) error {
	return database.DB.Delete(&models.GameStatistics{}, "id = ?", id).Error
}
