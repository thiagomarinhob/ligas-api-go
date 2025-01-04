package routes

import (
	"go-api-ligas/controllers"
	"go-api-ligas/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rotas públicas
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)
	router.POST("/refresh", controllers.RefreshToken)

	// Grupo de rotas protegidas
	auth := router.Group("/")
	auth.Use(middlewares.Auth())
	{
		// Rotas de produtos

		leagueGroup := auth.Group("/leagues")
		{
			leagueGroup.POST("/", controllers.CreateLeague)
			leagueGroup.GET("/", controllers.GetLeagues)
			leagueGroup.GET("/:id", controllers.GetLeagueByID)
			leagueGroup.GET("/:id/standings", controllers.GetLeagueStandings)
			leagueGroup.PUT("/:id", controllers.UpdateLeague)
			leagueGroup.DELETE("/:id", controllers.DeleteLeague)
		}

		teamGroup := auth.Group("/teams")
		{
			teamGroup.POST("/", controllers.CreateTeam)
			teamGroup.GET("/", controllers.GetTeams)
			teamGroup.GET("/:id", controllers.GetTeamByID)
			teamGroup.PUT("/:id", controllers.UpdateTeam)
			teamGroup.DELETE("/:id", controllers.DeleteTeam)
		}

		playerGroup := auth.Group("/players")
		{
			playerGroup.POST("/", controllers.CreatePlayer)
			playerGroup.GET("/", controllers.GetPlayers)
			playerGroup.GET("/:id", controllers.GetPlayerByID)
			playerGroup.PUT("/:id", controllers.UpdatePlayer)
			playerGroup.DELETE("/:id", controllers.DeletePlayer)
		}

		gameGroup := auth.Group("/games")
		{
			gameGroup.POST("/", controllers.CreateGame)
			gameGroup.GET("/", controllers.GetGames)
			gameGroup.GET("/:id", controllers.GetGameByID)
			gameGroup.PUT("/:id", controllers.UpdateGame)
			gameGroup.DELETE("/:id", controllers.DeleteGame)
		}

		statsGroup := auth.Group("/game-statistics")
		{
			statsGroup.POST("/", controllers.CreateGameStatistics)
			statsGroup.GET("/", controllers.GetGameStatistics)
			statsGroup.GET("/:id", controllers.GetGameStatisticsByID)
			statsGroup.PUT("/:id", controllers.UpdateGameStatistics)
			statsGroup.DELETE("/:id", controllers.DeleteGameStatistics)
		}

	}
}
