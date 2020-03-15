package main

import (
	// "github.com/YosukeCSato/go_sample/routes"

	"m/routes"
	"m/sessions"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	store := sessions.NewDummyStore()
	router.Use(sessions.StartDefaultSession(store))

	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogIn)
	}

	router.GET("/", routes.Home)
	router.GET("/login", routes.LogIn)
	router.GET("/signup", routes.SignUp)
	router.GET("/loggedin", routes.LoggedIn)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")

}
