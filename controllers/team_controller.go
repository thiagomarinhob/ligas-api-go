package controllers

import (
	"go-api-ligas/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context) {
	var input struct {
		Name           string  `json:"name"`
		Acronym        string  `json:"acronym"`
		LeagueID       string  `json:"league_id"`
		City           string  `json:"city"`
		State          string  `json:"state"`
		FoundationDate *string `json:"foundation_date"`
		Badge          string  `json:"badge"`
		Description    string  `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := services.CreateTeam(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

func GetTeams(c *gin.Context) {
	teams, err := services.GetTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func GetTeamByID(c *gin.Context) {
	id := c.Param("id")

	team, err := services.GetTeamByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := services.UpdateTeam(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func DeleteTeam(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteTeam(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team deleted successfully"})
}
