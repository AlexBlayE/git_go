package routes

import (
	"git_go/src/controllers"
	"git_go/src/middlewares"
	"git_go/src/schemas"

	"github.com/gin-gonic/gin"
)

func UserRouter(app *gin.Engine, groupRoute string) { // -> /user
	api := app.Group(groupRoute)

	api.POST("/create",
		middlewares.ValidateBody(new(schemas.CreateUserSchema)),
		controllers.CreateUserController,
	)
}
