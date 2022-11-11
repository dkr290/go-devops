package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-devops/go-gin-notes/controllers"
	"github.com/dkr290/go-devops/go-gin-notes/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	//loading static files
	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")
	models.ConnectDatabase()
	models.DbMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello gin",
		})
	})

	log.Println("Server is started!")
	r.Run("127.0.0.1:8080") // Default is 8080
}
