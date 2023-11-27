package middleware

import (
	"cardemo-crud/User/models"
	"cardemo-crud/helpers"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("JWT_SEC_KEY"))

type JwtClaim struct {
	Email string
	ID    int
	jwt.StandardClaims
}

func Login(r *gin.Context) {
	var user models.LoginUserRequest
	r.ShouldBindJSON(&user)
	varErr := user.LoginValidate()
	if varErr != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": varErr.Error()})
		r.Abort()
		return
	}
	userPassHashString, err := models.GetUserByEmail(user.Email)
	if err != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, err.Error())
		return
	}
	if userPassHashString == "" {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Please Use valid Email")
		return
	}
	passErr := helpers.VerifyHashedData(userPassHashString, user.Password)
	if passErr != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, passErr.Error())
		return
	}
	userData, _ := models.LoginUser(user.Email, userPassHashString)
	if userData.Password == "" && userData.Email == "" {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Unauthorized User")
		r.Abort()
		return
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JwtClaim{Email: userData.Email, ID: int(userData.ID), StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, err.Error())
		return
	}
	data := map[string]interface{}{"User_Data": token.Claims, "Tokan": tokenString}
	helpers.ApiSuccess(r, http.StatusOK, 1000, data)

}

func AuthMiddleware() gin.HandlerFunc {
	return func(r *gin.Context) {
		tokenString := r.GetHeader("Authorization")
		if tokenString == "" {
			helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Missing token")
			r.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			helpers.ApiFailure(r, http.StatusBadRequest, 1001, err.Error())
			r.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Invalid token claims")
			r.Abort()
			return
		}
		r.Set("user", claims)
		r.Next()
	}
}
