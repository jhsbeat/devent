// Package models handles models of MVC web application
package models

import (
	"time"

	"gorm.io/gorm"
)

// Event model
type Event struct {
	gorm.Model
	Id            uint      `gorm:"autoIncrement,unique,auto_increment"`
	Title         string    `gorm:"size:255,not null" form:"title" binding:"required"`
	Description   string    `form:"description" binding:"required"`
	StartAt       time.Time `gorm:"not null" form:"start_at" time_utc:"1" binding:"required"`
	EndAt         time.Time `gorm:"not null" form:"end_at" time_utc:"1" binding:"required"`
	LinkUrl       string    `gorm:"size:255,not null" form:"link_url" `
	FeaturedImage string    `gorm:"size:255"`
}

// GetEvents retrieves events data and set it up to the events slice variable.
func GetEvents(db *gorm.DB, events *[]Event) (err error) {
	err = db.Find(events).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateEvent created a event data.
func CreateEvent(db *gorm.DB, event *Event) (err error) {
	err = db.Create(event).Error
	if err != nil {
		return err
	}
	return nil
}
