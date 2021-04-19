package controllers

import (
	"devent/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEvents retrieves events data and reders `events/index.tmpl`.
func (repo *Repo) GetEvents(c *gin.Context) {
	var events []models.Event
	err := models.GetEvents(repo.Db, &events)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "events/index.tmpl", gin.H{
		"events": events,
	})
}

// NewEvent brings up an input form page for a new event registration (`events/new.tmpl`).
func (repo *Repo) NewEvent(c *gin.Context) {
	c.HTML(http.StatusOK, "events/new.tmpl", gin.H{})
}

// CreateEvent creates a record and redirect the page to root(`/`).
func (repo *Repo) CreateEvent(c *gin.Context) {

	var event models.Event

	if err := c.ShouldBind(&event); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	if err := models.CreateEvent(repo.Db, &event); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(http.StatusFound, "/")
}
