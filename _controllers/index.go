package _controllers

import (
	"fmt"
	"keepm/_models"
	util "keepm/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AppIsDown handles requests when the application is down.
func AppIsDown(c *gin.Context) {
	fmt.Println("Application is down")
	c.HTML(http.StatusInternalServerError, "dbdown.html", gin.H{"content": "Application is down"})
}

// IndexGet handles GET requests for the dashboard page.
func IndexGet(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(util.Username)
	account := session.Get(util.AccountName)
	contact := new(_models.Contact)
	contacts, _ := contact.GetAll(1)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{"content": "Dashboard...",
		"account":  account,
		"user":     user,
		"contacts": contacts})
}
