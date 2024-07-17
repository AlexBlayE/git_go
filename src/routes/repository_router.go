package routes

import (
	"github.com/gin-gonic/gin"
)

func RepositoryRouter(app *gin.Engine, groupRoute string) { // -> /admin
	app.POST("/info") // permetrá verure el espai de tot, tots els repos etc

	app.POST("/settings") // Per exemple pot habilitar si es poden crear usuaris o no

}
