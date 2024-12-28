package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameStatistics struct {
	ID            string  `json:"id" gorm:"type:uuid;primaryKey"`
	GameID        string  `json:"game_id" gorm:"not null"` // Foreign key for Game
	Game          Game    `json:"game" gorm:"foreignKey:GameID;constrastring:OnDelete:CASCADE;"`
	PlayerID      string  `json:"player_id" gorm:"not null"` // Foreign key for Player
	Player        Player  `json:"player" gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE;"`
	Points        int     `json:"points" gorm:"default:0"`
	Rebounds      int     `json:"rebounds" gorm:"default:0"`
	Assists       int     `json:"assists" gorm:"default:0"`
	Steals        int     `json:"steals" gorm:"default:0"`
	Blocks        int     `json:"blocks" gorm:"default:0"`
	Fouls         int     `json:"fouls" gorm:"default:0"`
	MinutesPlayed float32 `json:"minutes_played" gorm:"type:real;default:0"`
}

func (gameStatistics *GameStatistics) BeforeCreate(tx *gorm.DB) (err error) {
	if gameStatistics.ID == "" {
		gameStatistics.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
