package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Player represents a player entity.
type Player struct {
	ID        string     `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Number    int        `json:"number" gorm:"not null"`
	TeamID    string     `json:"team_id" gorm:"not null"` // Foreign key for Team
	Team      Team       `json:"team" gorm:"foreignKey:TeamID;constraint:OnDelete:CASCADE;"`
	Position  string     `json:"position" gorm:"type:varchar(30);not null"`
	Height    float32    `json:"height" gorm:"type:float"`
	Weight    float32    `json:"weight" gorm:"type:float"`
	BirthDate *time.Time `json:"birth_date" gorm:"type:date"`
	Photo     string     `json:"photo" gorm:"type:varchar(255)"`
	Active    bool       `json:"active" gorm:"default:true"`
}

func (player *Player) BeforeCreate(tx *gorm.DB) (err error) {
	if player.ID == "" {
		player.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
