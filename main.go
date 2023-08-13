package main

import (
	"gin/hello-world/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	loginController := &controllers.LoginController{}
	userboardController := &controllers.UserboardController{}

	r.LoadHTMLFiles(
		"templates/view-login.html",
		"templates/view-userboard.html",
		"templates/login/login-form.html",
		"templates/login/css-login.html",
		"templates/login/js-login.html",
		"templates/userboard/userboard.html",
		"templates/userboard/css-userboard.html",
		"templates/userboard/js-userboard.html",
	)

	// Middleware penanganan error global
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()
		c.Next()
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/login")
	})

	r.GET("/login", loginController.ShowLoginForm)

	r.POST("/login", loginController.HandleLogin)

	r.GET("/userboard", userboardController.Index)

	r.Run() // Menjalankan server pada port default (8080)
}
