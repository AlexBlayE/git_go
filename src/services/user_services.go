package services

import (
	"fmt"
	"git_go/src/models"
	"git_go/src/schemas"

	"github.com/google/uuid"
)

func CreateUser(name string, pass string) (string, error) {
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

	fmt.Println(jwtToken)

	if err != nil {
		return "", err
	}

	_, err = models.User.InsertOne(schemas.User{
		Name:     name,
		Password: encryptedPassword,
		Token:    token,
	})

	if err != nil {
		return "", err
	}

	return jwtToken, err
}

func IsUserLogged(name string, password string) bool {
	// user, err := models.User.FindOne(primitive.M{
	// 	"name":     name,
	// 	// "password": password,
	// })

	// if err != nil {
	// 	return false
	// }

	// fmt.Println(user)

	if name == "qqq" {
		return true
	}

	return false
}
