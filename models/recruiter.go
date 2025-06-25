package models

import "gorm.io/gorm"

type Recruiter struct {
	gorm.Model
	ID             uint   `json:"id" gorm:"primaryKey"`
	Name           string `gorm:"not null" json:"name"`
	Email          string `gorm:"unique;not null" json:"email"`
	Password       string `gorm:"not null" json:"password"`
	CompanyName    string `gorm:"not null" json:"company_name"`
	Role           string `json:"role"`
	CompanyAddress string `json:"company_address"`
	CompanyCity    string `json:"company_city"`
	Position       string `json:"position"`
	LinkedinURL    string `json:"linkedin_url"`
}
