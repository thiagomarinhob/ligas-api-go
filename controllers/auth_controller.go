package controllers

import (
	"go-api-ligas/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Perform user login
// @Description Authenticate a user using email and password, and return access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body struct {
// @     email string `json:"email" example:"user@example.com"`
// @     password string `json:"password" example:"password123"`
// @} true "User credentials"
// @Success 200 {object} map[string]interface{} "Tokens and user information"
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Router /login [post]

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o serviço para realizar o login
	response, err := services.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Generate a new access token using a valid refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param token body struct {
// @     refresh_token string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
// @} true "Refresh token"
// @Success 200 {object} map[string]string "New access token"
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Router /refresh-token [post]

func RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o serviço para renovar o token
	response, err := services.RefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Logout godoc
// @Summary Perform user logout
// @Description Revoke the refresh token for the authenticated user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param refresh_token body struct {
// @     refresh_token string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
// @} true "Refresh token"
// @Success 200 {object} gin.H{"message": "Logout realizado com sucesso"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /logout [post]

func Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o serviço para realizar o logout
	if err := services.Logout(userID.(string), input.RefreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}
