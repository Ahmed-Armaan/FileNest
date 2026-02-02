package utils

//import (
//	"errors"
//
//	"github.com/Ahmed-Armaan/FileNest/database"
//	"github.com/gin-gonic/gin"
//)
//
//func GetUserFromGoogleId(c *gin.Context) (*database.User, error) {
//	googleId, exist := c.Get("googleId")
//	if !exist {
//		return &database.User{}, errors.New("Could not resolve user")
//	}
//	googleIdStr, ok := googleId.(string)
//	if !ok {
//		return &database.User{}, errors.New("Could not resolve user")
//	}
//
//	user, err := database.GetUserByGoogleID(googleIdStr)
//	if err != nil {
//		return &database.User{}, errors.New("Database error")
//	}
//
//	return user, nil
//}
