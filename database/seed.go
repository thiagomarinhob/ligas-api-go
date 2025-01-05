package database

import (
	"fmt"
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

	teamsIDs := []string{}
	for i := 1; i <= 11; i++ {
		teamID := uuid.NewString()
		team := models.Team{
			ID:             teamID,
			Name:           fmt.Sprintf("Time %d", i),
			Acronym:        fmt.Sprintf("TM%d", i),
			LeagueID:       leagueID,
			City:           "Fortaleza",
			State:          "Ceará",
			FoundationDate: nil,
			Badge:          "",
			Description:    fmt.Sprintf("Time %d", i),
		}

		teamsIDs = append(teamsIDs, teamID)
		if err := db.FirstOrCreate(&team).Error; err != nil {
			return err
		}

		for j := 1; j <= 10; j++ {
			player := models.Player{
				ID:        uuid.NewString(),
				Name:      fmt.Sprintf("Jogador%dT%d", j, i),
				Number:    j,
				TeamID:    teamID,
				Position:  "Armador",
				Height:    1.80,
				Weight:    80,
				BirthDate: func(t time.Time) *time.Time { return &t }(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				Photo:     "",
			}

			if err := db.FirstOrCreate(&player).Error; err != nil {
				return err
			}
		}
	}

	// Cria os jogos entre as equipes
	gameIDs := []string{}
	for i := 0; i < len(teamsIDs); i++ {
		for j := i + 1; j < len(teamsIDs); j++ { // Evita duplicar jogos entre as mesmas equipes
			gameID := uuid.New().String()
			game := models.Game{
				ID:          gameID,
				LeagueID:    leagueID,
				TeamAID:     teamsIDs[i],
				TeamBID:     teamsIDs[j],
				DateTime:    time.Now().AddDate(0, 0, rand.Intn(30)), // Data aleatória nos próximos 30 dias
				Location:    "Estádio " + teamsIDs[i] + " x " + teamsIDs[j],
				PointsTeamA: rand.Intn(100), // Pontuação aleatória para Team A (0 a 9)
				PointsTeamB: rand.Intn(100), // Pontuação aleatória para Team B (0 a 9)
				Status:      "completed",
				Description: "Jogo gerado automaticamente.",
			}
			gameIDs = append(gameIDs, gameID)
			if err := db.FirstOrCreate(&game).Error; err != nil {
				return err
			}
		}
	}

	// Criação de estatísticas para cada jogador em cada jogo
	for _, gameID := range gameIDs {
		for _, teamID := range teamsIDs {
			var players []models.Player
			if err := db.Where("team_id = ?", teamID).Find(&players).Error; err != nil {
				return err
			}

			for _, player := range players {
				stats := models.GameStatistics{
					ID:            uuid.New().String(),
					GameID:        gameID,
					PlayerID:      player.ID,
					Points:        rand.Intn(15), // Valor fixo ou gerado aleatoriamente
					ThreePoints:   rand.Intn(10),
					Rebounds:      rand.Intn(7),
					Assists:       rand.Intn(10),
					Steals:        rand.Intn(6),
					Blocks:        rand.Intn(3),
					Fouls:         rand.Intn(5),
					MinutesPlayed: rand.Float32() * 40, // Minutos jogados (0 a 40)
				}
				if err := db.Create(&stats).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
