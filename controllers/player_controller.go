package controllers

import (
	"go-api-ligas/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePlayer(c *gin.Context) {
	var input struct {
		Name      string  `json:"name"`
		Number    int     `json:"number"`
		TeamID    string  `json:"team_id"`
		Position  string  `json:"position"`
		Height    float32 `json:"height"`
		Weight    float32 `json:"weight"`
		BirthDate *string `json:"birth_date"`
		Photo     string  `json:"photo"`
		Active    bool    `json:"active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := services.CreatePlayer(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}

func GetPlayers(c *gin.Context) {
	players, err := services.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func GetPlayerByID(c *gin.Context) {
	id := c.Param("id")

	player, err := services.GetPlayerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := services.UpdatePlayer(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func DeletePlayer(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeletePlayer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}
