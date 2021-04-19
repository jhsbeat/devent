package controllers

import (
	"devent/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 이벤트 목록 조회
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

// 이벤트 등록 화면
func (repo *Repo) NewEvent(c *gin.Context) {
	c.HTML(http.StatusOK, "events/new.tmpl", gin.H{})
}

// 이벤트 생성
func (repo *Repo) CreateEvent(c *gin.Context) {

	var event models.Event

	// TODO(hosang): 추후 conventional binding으로 수정
	var err error
	layout := "2006-01-02T15:04:05.000Z"
	event.Title = c.PostForm("title")
	event.Description = c.PostForm("description")
	event.StartAt, err = time.Parse(layout, c.PostForm("start_at"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	event.EndAt, err = time.Parse(layout, c.PostForm("end_at"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	event.LinkUrl = c.PostForm("link_url")

	err = models.CreateEvent(repo.Db, &event)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
