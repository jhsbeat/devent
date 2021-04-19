package controllers

import (
	"devent/database"
	"devent/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

func New() *Repo {
	db := database.InitDb()
	db.AutoMigrate(&models.Event{})
	return &Repo{Db: db}
}
