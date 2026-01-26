package handlers

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	user, err := utils.GetUserFromGoogleId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"user":    user.UserName,
		"profile": user.ProfileImage,
		"email":   user.Email,
	})
}
