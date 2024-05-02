package main

import (
	"fmt"
	"git_go/src/routes"
	"net/http"

	githttp "github.com/AaronO/go-git-http"
	"github.com/AaronO/go-git-http/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	go repo()

	func() {
		app := gin.Default()

		routes.UserRouter(app, "/api/user")

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
		// info.Repo
		// Disallow Pushes (making git server pull only)// TODO: en un futur fer aixó per depen dels usuaris
		if info.Push {
			return false, nil
		}

		fmt.Println(info.Repo)
		fmt.Println(info.Username)

		// Typically this would be a database lookup
		// if services.IsUserLogged(info.Username, info.Password) { // TODO: fer que miri si existeix l'usuari, si conincideix la contraseña y si te els permisos per aquest repo
		// 	return true, nil
		// }

		// return false, nil
		return true, nil
	})

	http.Handle("/", authenticator(repos)) //

	http.ListenAndServe(":7000", nil) //

	// Fer una api amb gin que permieti llistar esl repos per usuaris i admins
}
