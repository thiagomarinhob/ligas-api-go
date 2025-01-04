package database

import (
	"fmt"
	"log"

	"go-api-ligas/config"
	"go-api-ligas/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migrar os modelos
	DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.League{},
		&models.Team{},
		&models.Player{},
		&models.Game{},
		&models.GameStatistics{},
	)

	// if err := SeedAll(DB); err != nil {
	// 	log.Fatalf("Erro ao executar seed: %v", err)
	// }

	// if err := SeedLigas(DB); err != nil {
	// 	log.Fatalf("Erro ao executar seed: %v", err)
	// }

	// if err := SeedTeams(DB); err != nil {
	// 	log.Fatalf("Erro ao executar seed: %v", err)
	// }

}
