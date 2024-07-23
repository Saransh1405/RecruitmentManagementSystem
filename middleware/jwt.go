package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("FIRSTASSIGNMENT")

func GenerateToken(userID uint, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"user_type": userType,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
