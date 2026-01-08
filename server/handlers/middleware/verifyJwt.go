package middleware

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookies, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "No authention string found",
			})
			return
		}

		claims, err := utils.VerifyJwt(cookies)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "token expired",
				})
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Authentication information unavailable",
			})
			return
		}

		googleId, err := claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Authentication information invalid",
			})
			return
		}

		user, err := database.GetUserByGoogleID(googleId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user",
			})
			return
		}

		c.Set("username", user.UserName)
		c.Set("profile", user.ProfileImage)
		c.Set("email", user.Email)
		c.Next()
	}
}
