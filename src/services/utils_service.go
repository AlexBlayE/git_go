package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func IsValidPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func CreateJwt(token string) (string, error) {
	creator := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": token,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	str, err := creator.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return str, err
}

func ValidateJwt(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT_SECRET"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["token"].(string), nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
