package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"-"`
	Role         string `json:"role"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Education    string `json:"education"`
	Skills       string `json:"skills"`
	LinkedinURL  string `json:"linkedin_url"`
	GithubURL    string `json:"github_url"`
	PortfolioURL string `json:"portfolio_url"`
}
