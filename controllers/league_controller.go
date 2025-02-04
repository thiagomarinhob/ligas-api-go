package controllers

import (
	"go-api-ligas/services"
	"net/http"
	"strconv"

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

	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	// Associar o userID ao input
	input.UserID = userID

	league, err := services.CreateLeague(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, league)
}

func GetLeagues(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	leagues, err := services.GetLeagues(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leagues)
}

func GetLeagueByID(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	id := c.Param("id")

	league, err := services.GetLeagueByID(userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

func UpdateLeague(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	id := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	league, err := services.UpdateLeague(userID, id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

func DeleteLeague(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	id := c.Param("id")

	if err := services.DeleteLeague(userID, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
}

func GetLeagueStandings(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	id := c.Param("id")

	standings, err := services.GetLeagueStandings(userID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, standings)
}

func GetTotalPointsRanking(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	id := c.Param("id")
	limitParam := c.DefaultQuery("limit", "10") // Usar query parameter para o limite

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	ranking, err := services.GetTotalPointsRanking(userID, id, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ranking)
}

func GetTotalThreePointsRanking(c *gin.Context) {
	// Extrair o userID do header
	userID := c.GetHeader("User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User-ID header is required"})
		return
	}

	id := c.Param("id")
	limitParam := c.DefaultQuery("limit", "10") // Usar query parameter para o limite

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	ranking, err := services.GetTotalThreePointsRanking(userID, id, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ranking)
}
