package _controllers

import (
	"fmt"
	"keepm/_models"
	"keepm/globals"
	util "keepm/util"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginGet handles GET requests for the login page.
func LoginGet(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(util.Username)

	if globals.AppDown {
		AppIsDown(c)
		return
	}

	if user != nil {
		c.Redirect(http.StatusSeeOther, "/") // Redirect to home if already logged in.
	} else {
		c.HTML(http.StatusOK, "_login.html", gin.H{"content": "", "user": user})
	}
}

// LoginPost handles POST requests for user login.
func LoginPost(c *gin.Context) {
	fmt.Println("LoginPost")
	session := sessions.Default(c)
	user := session.Get(util.Username)
	if user != nil {
		c.HTML(http.StatusBadRequest, "_login.html", gin.H{"content": "Please logout first"})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if util.EmptyUserPass(username, password) {
		c.HTML(http.StatusBadRequest, "_login.html", gin.H{"content": "Parameters can't be empty"})
	}

	u := new(_models.User)
	authenticatedUser, err := u.Authenticate(username, password)
	globals.LOG("authenticatedUser: ", authenticatedUser)
	if err != nil {
		log.Println("Error authenticating user:", err)
		c.Redirect(http.StatusSeeOther, "/")
		//c.HTML(http.StatusInternalServerError, "dbdown.html", gin.H{"content": "Application is down"})
		return
	}

	if authenticatedUser != nil {
		session.Set(util.Username, username)

		a := new(_models.Account)
		account, _ := a.Get(authenticatedUser.AccountID)
		session.Set(util.AccountID, account.ID)
		session.Set(util.AccountName, account.Name)

		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "_login.html", gin.H{"content": "Failed to save session"})
		}
		globals.LOG("util.Username: ", session.Get(util.Username))
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(http.StatusUnauthorized, "_login.html", gin.H{"content": "Incorrect username or password"})
	}
}

// LogoutGet handles GET requests for user logout.
func LogoutGet(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(util.Username)
	if user == nil {
		log.Println("Invalid session token")
		return
	}
	session.Delete(util.Username)
	if err := session.Save(); err != nil {
		log.Println("Failed to save session:", err)
	}
	c.Redirect(http.StatusSeeOther, "/")
}
