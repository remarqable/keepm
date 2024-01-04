package _controllers

import (
	"keepm/globals"
	util "keepm/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired is a middleware function that checks if a user is logged in.
// It redirects to the login page if the user is not authenticated.
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(util.Username)
	if user == nil {
		globals.LOG("User not logged in")

		// Redirect to login page.
		c.Redirect(http.StatusMovedPermanently, "/login")

		// Abort the current request flow.
		c.Abort()
		return
	}
	c.Next() // Proceed to the next handler if the user is authenticated.
}
