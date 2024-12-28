package controllers

import (
	"go-api-ligas/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGame godoc
// @Summary Create a new game
// @Description Add a new game to a league, specifying teams, date, and other details
// @Tags Games
// @Accept json
// @Produce json
// @Param game body struct {
// @     league_id string `json:"league_id" example:"123"`
// @     team_a_id string `json:"team_a_id" example:"456"`
// @     team_b_id string `json:"team_b_id" example:"789"`
// @     date_time string `json:"date_time" example:"2024-11-22T15:00:00"`
// @     location string `json:"location" example:"Stadium A"`
// @     points_team_a int `json:"points_team_a" example:"0"`
// @     points_team_b int `json:"points_team_b" example:"0"`
// @     status string `json:"status" example:"scheduled"`
// @     description string `json:"description" example:"Championship Final"`
// @} true "Game details"
// @Success 201 {object} models.Game
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /games [post]

func CreateGame(c *gin.Context) {
	var input struct {
		LeagueID    string `json:"league_id"`
		TeamAID     string `json:"team_a_id"`
		TeamBID     string `json:"team_b_id"`
		DateTime    string `json:"date_time"`
		Location    string `json:"location"`
		PointsTeamA int    `json:"points_team_a"`
		PointsTeamB int    `json:"points_team_b"`
		Status      string `json:"status"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := services.CreateGame(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

// GetGames godoc
// @Summary List all games
// @Description Retrieve a list of all games, including details about leagues and teams
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {array} models.Game
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /games [get]

func GetGames(c *gin.Context) {
	games, err := services.GetGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

// GetGameByID godoc
// @Summary Get game details
// @Description Retrieve details of a specific game by its ID
// @Tags Games
// @Accept json
// @Produce json
// @Param id path string true "Game ID"
// @Success 200 {object} models.Game
// @Failure 404 {object} gin.H{"error": "Game not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /games/{id} [get]

func GetGameByID(c *gin.Context) {
	id := c.Param("id")

	game, err := services.GetGameByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

// UpdateGame godoc
// @Summary Update an existing game
// @Description Update details of a specific game, such as scores, location, or status
// @Tags Games
// @Accept json
// @Produce json
// @Param id path string true "Game ID"
// @Param game body map[string]interface{} true "Updated game details"
// @Success 200 {object} models.Game
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /games/{id} [put]

func UpdateGame(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := services.UpdateGame(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

// DeleteGame godoc
// @Summary Delete a game
// @Description Remove a game by its ID
// @Tags Games
// @Accept json
// @Produce json
// @Param id path string true "Game ID"
// @Success 200 {object} gin.H{"message": "Game deleted successfully"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /games/{id} [delete]

func DeleteGame(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteGame(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}
