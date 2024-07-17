package routes

import (
	"git_go/src/controllers"
	"git_go/src/middlewares"
	"git_go/src/schemas"

	"github.com/gin-gonic/gin"
)

func UserRouter(app *gin.Engine, groupRoute string) { // -> /user
	api := app.Group(groupRoute)

	api.POST("/register",
		middlewares.ValidateBody(new(schemas.CreateUserSchema)),
		controllers.CreateUserController,
	)

	api.POST("/login",
		middlewares.ValidateBody(new(schemas.GetTokenSchema)),
		controllers.GetTokenController,
	)

	api.GET("/info",
		middlewares.AuthMiddleware,
		controllers.GetInfo,
	)

	// TODO: crear grup si es tenen permisos
	// TODO: fer lo del permisos i admin etc

}
