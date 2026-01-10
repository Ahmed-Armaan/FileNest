package handlers

import (
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	// JWT authentication middleware sets these data
	// since the middleware hit the DB for user account authentication and get user details as a result.
	// I decided to use the fetched data rather than making new request.
	userName, _ := c.Get("username")
	profile, _ := c.Get("profile")
	email, _ := c.Get("email")

	c.JSON(200, gin.H{
		"user":    userName,
		"profile": profile,
		"email":   email,
	})
}
