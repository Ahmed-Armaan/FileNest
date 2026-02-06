package handlers

import (
	"net/http"
	"os"

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

func Test_user(c *gin.Context) {
	frontendUrl := os.Getenv("FRONTEND_URI")
	googleId := os.Getenv("TEST_GOOGLE_ID")
	jwtToken, err := utils.SignJwt(googleId)
	if err != nil {
		c.Redirect(302, frontendUrl+"/?error=response_construction_error")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session",
		Value:    jwtToken,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.Redirect(303, frontendUrl+"/home")
}
