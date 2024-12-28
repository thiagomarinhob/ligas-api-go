package services

import (
	"errors"
	"go-api-ligas/repository"
	"go-api-ligas/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Login realiza o login do usuário e gera tokens
func Login(email, password string) (map[string]interface{}, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	// Verifica a senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	// Gera tokens
	accessToken, err := utils.GenerateToken(user.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	// Salva o refresh token no banco de dados
	if err := repository.SaveRefreshToken(user.ID, refreshToken, time.Hour*24*7); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// RefreshToken gera um novo Access Token com base no Refresh Token
func RefreshToken(refreshToken string) (map[string]interface{}, error) {
	storedToken, err := repository.GetRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("refresh token inválido")
	}

	// Verifica se expirou
	if storedToken.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("refresh token expirado")
	}

	// Gera novo Access Token
	accessToken, err := utils.GenerateToken(storedToken.UserID, time.Minute*15)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"access_token": accessToken,
	}, nil
}

// Logout revoga o refresh token do usuário
func Logout(userID, refreshToken string) error {
	return repository.DeleteRefreshToken(userID, refreshToken)
}
