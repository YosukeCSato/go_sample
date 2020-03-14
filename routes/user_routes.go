package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup")
	username := ctx.PostForm("username")
	email := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	println("username: " + username)
	println("email: " + email)
	println("password: " + password)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/loggedin")
}

func UserLogIn(ctx *gin.Context) {
	println("post/login")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	println("username: " + username)
	println("password: " + password)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/loggedin")
}
