// Package controllers handles controllers of MVC web application
package controllers

import (
	"devent/database"
	"devent/models"

	"gorm.io/gorm"
)

// Repo is a struct type contains Gorm DB pointer.
type Repo struct {
	Db *gorm.DB
}

// New function initialize database,runs migrations and return the reference of Repo.
func New() *Repo {
	db := database.InitDb()
	db.AutoMigrate(&models.Event{})
	return &Repo{Db: db}
}
