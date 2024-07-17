package controllers

import (
	"fmt"
	"git_go/src/errors"
	"git_go/src/schemas"
	"git_go/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserController(ctx *gin.Context) {
	val, exists := ctx.Get("body")
	fmt.Println("a")

	if !exists {
		ctx.JSON(http.StatusBadRequest, errors.UserNotCreated)
		return
	}

	body := val.(*schemas.CreateUserSchema)

	token, err := services.CreateUser(body.Name, body.Password, body.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.UserNotCreated)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func GetTokenController(ctx *gin.Context) {
	val, exists := ctx.Get("body")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, errors.CantGetUserTokn)
		return
	}

	body := val.(*schemas.GetTokenSchema)

	token, isValid := services.RefreshToken(body.Email, body.Password)

	if !isValid {
		ctx.JSON(http.StatusUnauthorized, errors.CantGetUserTokn)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetInfo(ctx *gin.Context) {

}
