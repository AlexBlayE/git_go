package services

import (
	"fmt"
	"git_go/src/models"
	"git_go/src/schemas"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(name string, pass string, email string) (string, error) {
	encryptedPassword, err := EncryptPassword(pass)

	if err != nil {
		return "", err
	}

	uuidToken, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	token := uuidToken.String()

	jwtToken, err := CreateJwt(token)

	if err != nil {
		return "", err
	}

	_, err = models.User.InsertOne(schemas.User{
		Email:    email,
		Name:     name,
		Password: encryptedPassword,
		Token:    token,
	})

	if err != nil {
		return "", err
	}

	return jwtToken, err
}

func RefreshToken(email string, password string) (string, bool) {
	user, err := models.User.FindOne(primitive.M{
		"email": email,
	})

	if err != nil {
		return "", false
	}

	err = IsValidPassword(password, user.Password)

	if err != nil {
		return "", false
	}

	newToken, err := uuid.NewRandom()

	if err != nil {
		return "", false
	}

	_, err = models.User.UpdateOne(
		primitive.M{"email": email},
		primitive.M{"token": newToken.String()},
	)

	if err != nil {
		fmt.Println(err)
		return "", false
	}

	jwt, err := CreateJwt(newToken.String())

	if err != nil {
		return "", false
	}

	return jwt, true
}

func IsAuthTokenExists(token string) bool {
	_, err := models.User.FindOne(primitive.M{
		"token": token,
	})

	return err != nil
}
