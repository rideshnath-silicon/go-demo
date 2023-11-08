package middleware

import (
	"fmt"
	"hmvcstructure/helper"
	"hmvcstructure/user/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("HMVCC")

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var user models.User
		bodyData := helper.RequestBody(w, r, user)
		err := loginValidate(bodyData)
		if err != nil {
			helper.ApiFailure(w, 101, err.Error())
			return
		}
		userData := models.LoginUser(bodyData.Email, bodyData.Password)
		// fmt.Println(userData)
		if userData.Password == "" && userData.Email == "" {
			helper.ApiFailure(w, 5001)
			return
		}
		expirationTime := time.Now().Add(1 * time.Hour)
		claims := &JwtClaim{Email: bodyData.Email, StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"token": "%s"}`, tokenString)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			helper.ApiFailure(w, 5002)
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
			helper.ApiFailure(w, 5001, err.Error())
			return
		}
		next(w, r)
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
