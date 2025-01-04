package database

import (
	"fmt"
	"log"
	"math/rand"
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

func SeedAll(db *gorm.DB) error {

	// userId := uuid.NewString()
	// user := models.User{
	// 	ID:       userId,
	// 	Name:     "admin",
	// 	Email:    "admin@admin.com",
	// 	Password: "admin",
	// }

	// if err := db.FirstOrCreate(&user).Error; err != nil {
	// 	return err
	// }

	leagueID := uuid.NewString()
	league := models.League{
		ID:          leagueID,
		Name:        "Liga Cearense de Basquete",
		Category:    "Adulto",
		Gender:      "M",
		StartDate:   time.Now(),
		EndDate:     nil,
		MaxPlayers:  20,
		MaxTeams:    10,
		Status:      "active",
		Description: "Liga Cearense de Basquete 2025",
		UserID:      "ae271ce7-0dff-4f12-bcef-94d3b640bbe3",
	}

	if err := db.FirstOrCreate(&league).Error; err != nil {
		return err
	}

	teams := []models.Team{
		{
			ID:             uuid.New().String(),
			Name:           "ABA Apuiares",
			Acronym:        "ABA",
			LeagueID:       leagueID,
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
			LeagueID:       leagueID,
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
			LeagueID:       leagueID,
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
			LeagueID:       leagueID,
			City:           "Kalangos",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "Kalangos",
		},
		{
			ID:             uuid.New().String(),
			Name:           "DMB",
			Acronym:        "DMB",
			LeagueID:       leagueID,
			City:           "Fortaleza",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "Kalangos",
		},
		{
			ID:             uuid.New().String(),
			Name:           "Henrique Jorge",
			Acronym:        "HJB",
			LeagueID:       leagueID,
			City:           "Fortaleza",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "Henrique Jorge",
		},
		{
			ID:             uuid.New().String(),
			Name:           "Info Digital",
			Acronym:        "IFD",
			LeagueID:       leagueID,
			City:           "Rio Grande do Norte",
			State:          "RN",
			FoundationDate: nil,
			Badge:          "",
			Description:    "Info Digital",
		},
		{
			ID:             uuid.New().String(),
			Name:           "5 shots",
			Acronym:        "5SH",
			LeagueID:       leagueID,
			City:           "Fortaleza",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    "5 Shots",
		},
	}

	for _, team := range teams {
		// Adiciona a liga apenas se ela não existir no banco
		if err := db.FirstOrCreate(&team, models.League{ID: team.ID}).Error; err != nil {
			return err
		}
	}

	// Criar jogos
	// rand.Seed(time.Now().UnixNano()) // Inicializa a seed para gerar números aleatórios
	games := []models.Game{}
	// leagueID := teams[0].LeagueID // Assumindo que todos os times pertencem à mesma liga

	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ { // Evita duplicar jogos entre as mesmas equipes
			game := models.Game{
				ID:          uuid.NewString(),
				LeagueID:    leagueID,
				TeamAID:     teams[i].ID,
				TeamBID:     teams[j].ID,
				DateTime:    time.Now().AddDate(0, 0, rand.Intn(30)), // Data aleatória nos próximos 30 dias
				Location:    "Estádio " + teams[i].Name + " x " + teams[j].Name,
				PointsTeamA: rand.Intn(10), // Pontuação aleatória para Team A (0 a 9)
				PointsTeamB: rand.Intn(10), // Pontuação aleatória para Team B (0 a 9)
				Status:      "completed",
				Description: "Jogo gerado automaticamente.",
			}
			games = append(games, game)
			// fmt.Println("Foram criados jogos entre os times.", game)
		}
	}

	fmt.Println("Jogos que seram salvos.", len(games))

	for _, game := range games {
		if err := db.FirstOrCreate(&game).Error; err != nil {
			return err
		}
	}

	log.Printf("Foram criados %d jogos entre os times.", len(games))

	return nil
}
