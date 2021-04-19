// Controllers package to handle controllers of MVC web application
package controllers

import (
	"devent/database"
	"devent/models"

	"gorm.io/gorm"
)

// Repository type of GORM DB.
type Repo struct {
	Db *gorm.DB
}

func New() *Repo {
	db := database.InitDb()
	db.AutoMigrate(&models.Event{})
	return &Repo{Db: db}
}
