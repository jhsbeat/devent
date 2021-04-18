package controllers

import (
	"devent/database"
	"devent/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventRepo struct {
	Db *gorm.DB
}

func New() *EventRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Event{})
	return &EventRepo{Db: db}
}

// 이벤트 목록 조회
func (repository *EventRepo) GetEvents(c *gin.Context) {
	var events []models.Event
	err := models.GetEvents(repository.Db, &events)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "events/index.tmpl", gin.H{
		"events": events,
	})
}
