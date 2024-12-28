package repository

import (
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
