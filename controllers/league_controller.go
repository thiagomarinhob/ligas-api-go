package controllers

import (
	"go-api-ligas/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLeague(c *gin.Context) {
	var input struct {
		Name        string  `json:"name"`
		Category    string  `json:"category"`
		Gender      string  `json:"gender"`
		StartDate   string  `json:"start_date"`
		EndDate     *string `json:"end_date"`
		MaxPlayers  int     `json:"max_players"`
		MaxTeams    int     `json:"max_teams"`
		Status      string  `json:"status"`
		Description string  `json:"description"`
		UserID      string  `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	league, err := services.CreateLeague(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, league)
}

func GetLeagues(c *gin.Context) {
	leagues, err := services.GetLeagues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leagues)
}

func GetLeagueByID(c *gin.Context) {
	id := c.Param("id")

	league, err := services.GetLeagueByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

func UpdateLeague(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	league, err := services.UpdateLeague(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

func DeleteLeague(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteLeague(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
}
