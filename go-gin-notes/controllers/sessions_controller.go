package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home/login.html",
		nil,
	)
}

func SignupPage(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home/signup.html",
		nil,
	)
}
