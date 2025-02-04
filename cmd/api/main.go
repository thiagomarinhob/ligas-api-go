package main

import (
	"go-api-ligas/config"
	"go-api-ligas/database"
	"go-api-ligas/routes"

	_ "go-api-ligas/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"time"
)

// @title API Documentation
// @version 1.0
// @description This is a sample server for a sports league API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func main() {
	// Carregar as configurações
	config.LoadConfig()

	// Conectar ao banco de dados
	database.ConnectDatabase()

	// Iniciar o roteador
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"},                                  // Domínios permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                     // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Establishment-ID"}, // Cabeçalhos permitidos
		ExposeHeaders:    []string{"Content-Length"},                                              // Cabeçalhos expostos
		AllowCredentials: true,                                                                    // Permitir credenciais (cookies)
		MaxAge:           12 * time.Hour,                                                          // Tempo de cache do CORS
	}))

	// Configurar as rotas
	routes.SetupRoutes(server)

	// Swagger route
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor
	server.Run(":8080")
}
