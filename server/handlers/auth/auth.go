package auth

import (
	"net/http"
	"os"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
)

type TokenResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
	Scope                 string `json:"scope"`
}

type UserInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GetCredentials(c *gin.Context) {
	frontendUrl := os.Getenv("FRONTEND_URI")
	code := c.Query("code")
	oauthErr := c.Query("error")

	if oauthErr != "" {
		c.Redirect(302, frontendUrl+"/?error="+oauthErr)
		return
	}

	if code == "" {
		c.Redirect(302, frontendUrl+"/?error=Oauth_denied")
		return
	}

	tokenResponse, err := getTokens(code)
	if err != nil {
		c.Redirect(302, frontendUrl+"/?error=Oauth_denied")
		return
	}

	userInfo, err := getUserInfo(tokenResponse.AccessToken)
	if err != nil {
		c.Redirect(302, frontendUrl+"/?error=Oauth_denied")
		return
	}

	if err := database.InsertUser(userInfo.Name, userInfo.Sub, userInfo.Email, userInfo.Picture); err != nil {
		c.Redirect(http.StatusInternalServerError, frontendUrl+"/?error=database_error")
		return
	}

	jwtToken, err := utils.SignJwt(userInfo.Sub)
	if err != nil {
		c.Redirect(302, frontendUrl+"/?error=response_construction_error")
		return
	}

	c.SetCookie("session", jwtToken, 60*60*24*7, "", "", false, true)
	c.Redirect(303, frontendUrl+"/home")
}
