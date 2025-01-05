package controllers

import (
	"go-api-ligas/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGameStatistics(c *gin.Context) {
	var input struct {
		GameID        string  `json:"game_id"`
		PlayerID      string  `json:"player_id"`
		Points        int     `json:"points"`
		ThreePoints   int     `json:"three_points"`
		Rebounds      int     `json:"rebounds"`
		Assists       int     `json:"assists"`
		Steals        int     `json:"steals"`
		Blocks        int     `json:"blocks"`
		Fouls         int     `json:"fouls"`
		MinutesPlayed float32 `json:"minutes_played"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameStatistics, err := services.CreateGameStatistics(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gameStatistics)
}

func GetGameStatistics(c *gin.Context) {
	statistics, err := services.GetGameStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}

func GetGameStatisticsByID(c *gin.Context) {
	id := c.Param("id")

	statistics, err := services.GetGameStatisticsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}

func UpdateGameStatistics(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statistics, err := services.UpdateGameStatistics(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}

func DeleteGameStatistics(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteGameStatistics(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game statistics deleted successfully"})
}
