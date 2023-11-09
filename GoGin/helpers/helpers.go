package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ApiSuccess(c *gin.Context, status int, messageCode int, data interface{}) {
	type ApiSuccessResponse struct {
		Message string
		Success int
		Data    interface{}
	}
	message := Messagess(messageCode)
	Response := ApiSuccessResponse{
		Message: message,
		Success: 1,
		Data:    data,
	}
	c.JSON(status, Response)
}

func ApiFailure(c *gin.Context, status int, messageCode int, err string) {
	type ApiSuccessResponse struct {
		Message string
		Success int
		Error   string
	}
	message := Messagess(messageCode)
	Response := ApiSuccessResponse{
		Message: message,
		Success: 0,
		Error:   err,
	}
	c.JSON(status, Response)
}

func GetUserDataFromTokan(c *gin.Context) interface{} {
	userClaims, exists := c.Get("user")
	if !exists {
		return 0
	}
	// Access user data from claims
	userID := userClaims.(jwt.MapClaims)["Email"].(string)
	return userID
}
