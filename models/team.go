package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team represents a sports team entity.
type Team struct {
	ID             string     `json:"id" gorm:"type:uuid;primaryKey"`
	Name           string     `json:"name" gorm:"type:varchar(255);not null"`
	Acronym        string     `json:"acronym" gorm:"type:varchar(10);not null"`
	LeagueID       string     `json:"league_id" gorm:"not null"` // Foreign key for League
	League         League     `json:"league" gorm:"foreignKey:LeagueID;constraint:OnDelete:CASCADE;"`
	City           string     `json:"city" gorm:"type:varchar(100)"`
	State          string     `json:"state" gorm:"type:varchar(50)"`
	FoundationDate *time.Time `json:"foundation_date" gorm:"type:date"`
	Badge          string     `json:"badge" gorm:"type:varchar(255)"` //emblema
	Description    string     `json:"description" gorm:"type:text"`
}

func (team *Team) BeforeCreate(tx *gorm.DB) (err error) {
	if team.ID == "" {
		team.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
