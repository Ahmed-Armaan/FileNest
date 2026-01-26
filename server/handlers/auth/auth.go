package auth

import (
	"fmt"
	"net/http"
	"os"

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
		fmt.Printf("OAUTH_ERR: %s\n", oauthErr)
		c.Redirect(302, frontendUrl+"/?error="+oauthErr)
		return
	}

	if code == "" {
		fmt.Println("No code found")
		c.Redirect(302, frontendUrl+"/?error=Oauth_denied")
		return
	}

	tokenResponse, err := getTokens(code)
	if err != nil {
		fmt.Println("No token found")
		c.Redirect(302, frontendUrl+"/?error=Oauth_denied")
		return
	}

	userInfo, err := getUserInfo(tokenResponse.AccessToken)
	if err != nil {
		fmt.Println("No userInfo found")
		c.Redirect(302, frontendUrl+"/?error=Oauth_denied")
		return
	}

	jwtToken, err := utils.SignJwt(userInfo.Sub)
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
