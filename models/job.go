package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID            uint       `json:"id" gorm:"primaryKey"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Company       string     `json:"company"`
	Location      string     `json:"location"`
	ExpiredAt     *time.Time `json:"expired_at"`
	Status        string     `json:"status"`
	Address       string     `json:"address"`
	City          string     `json:"city"`
	Skills        string     `json:"skills"`
	Qualification string     `json:"qualification"`
}
