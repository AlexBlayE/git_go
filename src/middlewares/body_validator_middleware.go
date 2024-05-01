package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateBody(bodyType any) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val := validator.New()

		if err := ctx.BindJSON(bodyType); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := val.Struct(bodyType); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("body", bodyType)

		ctx.Next()
	}
}
