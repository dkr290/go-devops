package controllers

import (
	"net/http"

	"github.com/dkr290/go-devops/go-gin-notes/models"
	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {

	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"notes/new.html",
		nil,
	)

}

func NotesCreate(c *gin.Context) {

	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NoteCreate(name, content)

	c.Redirect(http.StatusMovedPermanently, "notes")

}
