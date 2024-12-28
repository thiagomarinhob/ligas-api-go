package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type League struct {
	ID          string     `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string     `json:"name" gorm:"type:varchar(255);not null"`
	Category    string     `json:"category" gorm:"type:varchar(50);not null"`
	Gender      string     `json:"gender" gorm:"type:varchar(1);not null"` // 'M', 'F', 'U'
	StartDate   time.Time  `json:"start_date" gorm:"type:date;not null"`
	EndDate     *time.Time `json:"end_date" gorm:"type:date"`
	MaxPlayers  int        `json:"max_players" gorm:"default:12"`
	MaxTeams    int        `json:"max_teams" gorm:"default:16"`
	Status      string     `json:"status" gorm:"type:varchar(20);default:'active'"` // 'active', 'closed', 'suspended'
	Description string     `json:"description" gorm:"type:text"`
	UserID      string     `json:"user_id" gorm:"not null"` // Foreign key for User
	User        User       `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (league *League) BeforeCreate(tx *gorm.DB) (err error) {
	if league.ID == "" {
		league.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
