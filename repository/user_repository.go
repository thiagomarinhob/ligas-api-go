package repository

import (
	"go-api-ligas/database"
	"go-api-ligas/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}
