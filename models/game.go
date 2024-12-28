package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey"`
	LeagueID    string    `json:"league_id" gorm:"not null"` // Foreign key for League
	League      League    `json:"league" gorm:"foreignKey:LeagueID;constraint:OnDelete:CASCADE;"`
	TeamAID     string    `json:"team_a_id" gorm:"not null"` // Foreign key for Team A
	TeamA       Team      `json:"team_a" gorm:"foreignKey:TeamAID;constraint:OnDelete:CASCADE;"`
	TeamBID     string    `json:"team_b_id" gorm:"not null"` // Foreign key for Team B
	TeamB       Team      `json:"team_b" gorm:"foreignKey:TeamBID;constraint:OnDelete:CASCADE;"`
	DateTime    time.Time `json:"date_time" gorm:"type:timestamp;not null"`
	Location    string    `json:"location" gorm:"type:varchar(255)"`
	PointsTeamA int       `json:"points_team_a" gorm:"default:0"`
	PointsTeamB int       `json:"points_team_b" gorm:"default:0"`
	Status      string    `json:"status" gorm:"type:varchar(20);default:'scheduled'"` //('scheduled', 'in progress', 'completed', 'canceled')
	Description string    `json:"description" gorm:"type:text"`
}

func (game *Game) BeforeCreate(tx *gorm.DB) (err error) {
	if game.ID == "" {
		game.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
