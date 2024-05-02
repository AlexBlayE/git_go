package controllers

import (
	"git_go/src/schemas"
	"git_go/src/services"

	"github.com/gin-gonic/gin"
)

func CreateUserController(ctx *gin.Context) {
	val, exists := ctx.Get("body")

	if !exists {
		ctx.Status(400)
		return
	}

	body := val.(*schemas.CreateUserSchema)

	token, err := services.CreateUser(body.Name, body.Password)

	if err != nil {
		ctx.JSON(400, gin.H{
			"err": "user not created",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"token": token,
	})
}
