package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`

	JobID uint `json:"job_id"`
	Job   Job  `json:"job" gorm:"foreignKey:JobID"`

	Status string `json:"status"`
}
