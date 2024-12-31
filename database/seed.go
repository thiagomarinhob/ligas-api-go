package database

import (
	"time"

	"go-api-ligas/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedLigas(db *gorm.DB) error {
	ligas := []models.League{
		{
			ID:          uuid.New().String(),
			Name:        "Liga Brasileira de Futebol",
			Category:    "Profissional",
			Gender:      "M",
			StartDate:   time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC),
			MaxPlayers:  11,
			MaxTeams:    20,
			Status:      "active",
			Description: "A principal liga nacional de futebol.",
			UserID:      "b0205fac-cee4-4297-b253-4db8b967352d", // Substitua por IDs válidos
		},
		{
			ID:          uuid.New().String(),
			Name:        "Liga Nacional de Basquete Feminino",
			Category:    "Adulto",
			Gender:      "F",
			StartDate:   time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC),
			MaxPlayers:  12,
			MaxTeams:    16,
			Status:      "active",
			Description: "Uma liga de destaque no cenário do basquete feminino.",
			UserID:      "b0205fac-cee4-4297-b253-4db8b967352d", // Substitua por IDs válidos
		},
	}

	for _, liga := range ligas {
		// Adiciona a liga apenas se ela não existir no banco
		if err := db.FirstOrCreate(&liga, models.League{ID: liga.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedTeams(db *gorm.DB) error {
	teams := []models.Team{
		{
			ID:             uuid.New().String(),
			Name:           "ABA Apuiares",
			Acronym:        "ABA",
			LeagueID:       "01562fab-4395-4532-9b38-4fa2dda8b145",
			City:           "Apuiarés",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "ABA Apuiarés",
		},
		{
			ID:             uuid.New().String(),
			Name:           "São Goncalo do Amarante",
			Acronym:        "SGA",
			LeagueID:       "01562fab-4395-4532-9b38-4fa2dda8b145",
			City:           "São Goncalo do Amarante",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "São Goncalo do Amarante",
		},
		{
			ID:             uuid.New().String(),
			Name:           "Universidade Federal do Ceará",
			Acronym:        "UFC",
			LeagueID:       "01562fab-4395-4532-9b38-4fa2dda8b145",
			City:           "Universidade Federal do Ceará",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "Universidade Federal do Ceará",
		},
		{
			ID:             uuid.New().String(),
			Name:           "Kalangos",
			Acronym:        "KLG",
			LeagueID:       "01562fab-4395-4532-9b38-4fa2dda8b145",
			City:           "Kalangos",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "Kalangos",
		},
	}

	for _, team := range teams {
		// Adiciona a liga apenas se ela não existir no banco
		if err := db.FirstOrCreate(&team, models.League{ID: team.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}
