package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello gin",
		})
	})

	log.Println("Server is started!")
	r.Run("127.0.0.1:8080") // Default is 8080
}
