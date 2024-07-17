package middlewares

import (
	"git_go/src/errors"
	"git_go/src/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, errors.NeedAuthHeaderBody)
		return
	}

	splitedHeader := strings.Split(authHeader, " ")

	if strings.ToLower(splitedHeader[0]) != "bearer" {
		ctx.JSON(http.StatusBadRequest, errors.NeedAuthHeaderBody)
		return
	}

	token, err := services.ValidateJwt(splitedHeader[1])

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.InvalidJWT)
		return
	}

	if !services.IsAuthTokenExists(token) {
		ctx.JSON(http.StatusBadRequest, errors.InvalidJWT)
		return
	}

	ctx.Next()
}
