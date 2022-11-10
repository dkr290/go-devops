package repository

import "gorm.io/gorm"

type Repository struct {
	DB     *gorm.DB
	DbUser string
	DbPass string
	DbPort string
	DbName string
	DbHost string
}

// creates new repository
func NewRepo(d *gorm.DB) *Repository {
	r := Repository{
		DB: d,
	}
	return &r
}
