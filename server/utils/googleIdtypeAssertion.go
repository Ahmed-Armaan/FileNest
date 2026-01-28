package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GoogleIdstring(c *gin.Context) (string, error) {
	googleId, exist := c.Get("googleId")
	if !exist {
		return "", errors.New("User resolution failed")
	}

	googleIdStr, ok := googleId.(string)
	if !ok {
		return "", errors.New("invalid googleId provided")
	}

	return googleIdStr, nil
}
