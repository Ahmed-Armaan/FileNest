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

	// insert new users into the DB
	user, err := database.GetUserByGoogleID(userInfo.Sub)
	if err != nil {
		user, err = database.InsertUser(userInfo.Name, userInfo.Sub, userInfo.Email, userInfo.Picture)
		if err != nil {
			c.Redirect(303, frontendUrl+"/?error=database_error")
			return
		}
	}

	//if err = database.InsertUser(userInfo.Name, userInfo.Sub, userInfo.Email, userInfo.Picture); err != nil {
	//	c.Redirect(303, frontendUrl+"/?error=database_error")
	//	return
	//}

	jwtToken, err := utils.SignJwt(userInfo.Sub)
	if err != nil {
		c.Redirect(302, frontendUrl+"/?error=response_construction_error")
		return
	}

	// get rootElementId
	rootNode, err := database.GetRootNodeId(user.ID)
	if err != nil {
		c.Redirect(302, frontendUrl+"/?error=root_node_not_found")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "rootNodeId",
		Value:    rootNode.ID.String(),
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "rootNodeUpdatedAt",
		Value:    rootNode.UpdatedAt.Format("2006-01-02 15:04:05"),
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	})

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
