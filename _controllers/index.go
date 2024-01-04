package _controllers

import (
	"fmt"
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
	fmt.Println("IndexGet")
	session := sessions.Default(c)
	user := session.Get(util.Username)
	account := session.Get(util.AccountName)
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"content": "Dashboard...", "account": account, "user": user})
}
