package handlers

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
)

func Me(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		googleIdStr, err := utils.GoogleIdstring(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		user, err := db.GetUserDataByGoogleId(googleIdStr, database.UserDbColums.UserName, database.UserDbColums.Email, database.UserDbColums.ProfileImage)
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
}
