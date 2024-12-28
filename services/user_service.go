// services/user_service.go
package services

import (
	"go-api-ligas/models"
	"go-api-ligas/repository"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser realiza o registro de um novo usuário
func RegisterUser(name, email, password string) (models.User, error) {
	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// Criação do modelo de usuário
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Chamada ao repositório para salvar no banco de dados
	if err := repository.CreateUser(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
