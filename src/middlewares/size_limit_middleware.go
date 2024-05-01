package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SizeLimitMiddleware(mgLimit int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.ContentLength > mgLimit {
			ctx.JSON(http.StatusRequestEntityTooLarge, "j") // TODO: cambiar el type
			return
		}

		ctx.Next()
	}
}
