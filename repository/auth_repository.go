package repository

import (
	"go-api-ligas/database"
	"go-api-ligas/models"
	"time"
)

// GetUserByEmail retorna um usuário com base no email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// SaveRefreshToken salva o refresh token no banco
func SaveRefreshToken(userID string, token string, duration time.Duration) error {
	refreshToken := models.Token{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(duration).Unix(),
	}
	return database.DB.Create(&refreshToken).Error
}

// GetRefreshToken retorna um refresh token com base no valor
func GetRefreshToken(token string) (models.Token, error) {
	var storedToken models.Token
	err := database.DB.Where("token = ?", token).First(&storedToken).Error
	return storedToken, err
}

// DeleteRefreshToken remove o refresh token do banco
func DeleteRefreshToken(userID, token string) error {
	return database.DB.Where("user_id = ? AND token = ?", userID, token).Delete(&models.Token{}).Error
}
