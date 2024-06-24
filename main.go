package main

import (
	"git_go/src/routes"
	"git_go/src/services"
	"net/http"

	githttp "github.com/AaronO/go-git-http"
	"github.com/AaronO/go-git-http/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

// TODO: en un futur fer que es puguin crear diversos rervidors y que es pugui redirigir la petició a un altre servidor si en aquest no está el repo

func main() {
	go repo()

	func() {
		app := gin.Default()

		routes.UserRouter(app, "/api/user")
		// Ruta de grups
		// Ruta de info de repositori
		// Ruta de admin i opcions

		// ruta per lo de git
		app.Use(func(ctx *gin.Context) {
			originalUrl := ctx.Request.URL
			originalUrl.Host = "localhost:7000"
			ctx.Redirect(308, originalUrl.String())
		})

		app.Run(":8080")
	}()
}

func repo() {
	repos := githttp.New("./repos")
	// Amb go-git crear la api per poder fer git merge a la main

	// TODO: em deixa clonarlo i no vui sense autentificarme
	authenticator := auth.Authenticator(func(info auth.AuthInfo) (bool, error) {
		return true, nil
		if services.IsUserLogged(info.Username, info.Password) {
			return true, nil
		}

		// TODO: aqui mirar si l'usuari te permisos pel repositori i quins permisos té
		// TODO: anar comprovant el tamañ del repositori per que no s'arribi al limit

		if info.Push {
			return false, nil
		}

		return false, nil
	})

	http.Handle("/", authenticator(repos)) //

	http.ListenAndServe(":7000", nil) //

	// Fer una api amb gin que permieti llistar esl repos per usuaris i admins
}
