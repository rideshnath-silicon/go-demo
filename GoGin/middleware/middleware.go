package middleware

import (
	"fmt"
	"gin/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("HMVCC")

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func Login(r *gin.Context) {
	var user models.User
	r.ShouldBindJSON(&user)
	err := loginValidate(user)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		r.Abort()
		return
	}
	userData := models.LoginUser(user.Email, user.Password)
	if userData.Password == "" && userData.Email == "" {
		r.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})
		r.Abort()
		return
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JwtClaim{Email: user.Email, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	r.JSON(http.StatusBadRequest, gin.H{"Tokan": tokenString})

}

func AuthMiddleware() gin.HandlerFunc {
	return func(r *gin.Context) {
		tokenString := r.GetHeader("Authorization")
		if tokenString == "" {
			r.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
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
			r.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized User"})
			r.Abort()
		}
		r.Next()
	}
}

func loginValidate(data models.User) error {
	if data.Email == "" {
		return fmt.Errorf("email is required")
	}
	if data.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
