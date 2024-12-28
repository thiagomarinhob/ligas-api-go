package repository

import (
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreateTeam(team *models.Team) error {
	return database.DB.Create(team).Error
}

func GetTeams() ([]models.Team, error) {
	var teams []models.Team
	err := database.DB.Preload("League").Find(&teams).Error
	return teams, err
}

func GetTeamByID(id string) (models.Team, error) {
	var team models.Team
	err := database.DB.Preload("League").First(&team, "id = ?", id).Error
	return team, err
}

func UpdateTeam(id string, updates map[string]interface{}) (models.Team, error) {
	var team models.Team
	if err := database.DB.Model(&team).Where("id = ?", id).Updates(updates).Error; err != nil {
		return models.Team{}, err
	}
	database.DB.First(&team, "id = ?", id)
	return team, nil
}

func DeleteTeam(id string) error {
	return database.DB.Delete(&models.Team{}, "id = ?", id).Error
}
