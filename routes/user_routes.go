package routes

import (
	"m/config"
	"m/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup")
	username := ctx.PostForm("username")
	email := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")

	db := config.DummyDB()
	if err := db.SaveUser(username, email, password); err != nil {
		println("Error: " + err.Error())
		ctx.HTML(http.StatusOK, "signup_failed.html", gin.H{})
		return
	}

	println("Signup success!")
	println("username: " + username)
	println("email: " + email)
	println("password: " + password)

	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: while loading user: " + err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	session := sessions.GetDefaultSession(ctx)
	session.Set("user", user)
	session.Save()
	println("session saved.")
	println("  sessionID: " + session.ID)
	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/loggedin")
}

func UserLogIn(ctx *gin.Context) {
	println("post/login")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db := config.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: " + err.Error())

		ctx.HTML(http.StatusOK, "login_failed.html", gin.H{})
		return
	}

	println("Authentication Success!")
	println("username: " + user.Username)
	println("email: " + user.Email)
	println("password: " + user.Password)
	session := sessions.GetDefaultSession(ctx)
	session.Set("user", user)
	session.Save()
	user.Authenticate()

	println("Session saved.")
	println("  sessionID: " + session.ID)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/loggedin")
}
