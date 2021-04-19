package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Id            uint   `gorm:"autoIncrement,unique,auto_increment"`
	Title         string `gorm:"size:255,not null"`
	Description   string
	StartAt       time.Time `gorm:"not null"`
	EndAt         time.Time `gorm:"not null"`
	LinkUrl       string    `gorm:"size:255,not null"`
	FeaturedImage string    `gorm:"size:255"`
}

// Get Events
func GetEvents(db *gorm.DB, event *[]Event) (err error) {
	err = db.Find(event).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateEvent(db *gorm.DB, event *Event) (err error) {
	err = db.Create(event).Error
	if err != nil {
		return err
	}
	return nil
}
