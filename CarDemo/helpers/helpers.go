package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func GetUserDataFromTokan(c *gin.Context) (map[string]interface{}, bool) {
	userClaims, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	// Access user data from claims
	userID := userClaims.(jwt.MapClaims)["ID"]
	userEmail := userClaims.(jwt.MapClaims)["Email"]
	response := map[string]interface{}{"Email": userEmail, "User_id": userID}

	return response, true
}

func HashData(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func VerifyHashedData(hashedString string, dataString string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(dataString))

	if err != nil {
		return err
	}
	return nil
}
