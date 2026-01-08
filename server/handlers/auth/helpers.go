package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func getTokens(code string) (*TokenResponse, error) {
	var tokenResponse TokenResponse
	reqUrl := os.Getenv("TOKEN_URI")

	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", os.Getenv("OAUTH_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("OAUTH_CLIENT_SECRET"))
	data.Set("redirect_uri", os.Getenv("REDIRECT_URI"))
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return &tokenResponse, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &tokenResponse, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return &tokenResponse, fmt.Errorf(
			"token exchange failed: %s | body: %s",
			res.Status,
			string(bodyBytes),
		)
	}

	err = json.NewDecoder(res.Body).Decode(&tokenResponse)
	if err != nil {
		return &tokenResponse, err
	}

	return &tokenResponse, nil
}

func getUserInfo(accessToken string) (*UserInfo, error) {
	reqUrl := "https://www.googleapis.com/oauth2/v3/userinfo"
	var googleUserInfo UserInfo

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return &googleUserInfo, err
	}

	authorizationStr := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Add("Authorization", authorizationStr)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &googleUserInfo, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &googleUserInfo, err
	}

	err = json.Unmarshal(body, &googleUserInfo)
	if err != nil {
		return &googleUserInfo, err
	}

	return &googleUserInfo, nil
}
