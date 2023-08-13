// controllers/login_controller.go
package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

// Di dalam file login_controller.go
func (ctrl *LoginController) ShowLoginForm(c *gin.Context) {
	// 1. Load template
	template, err := template.ParseFiles(
		"templates/view-login.html",
		"templates/login/css-login.html",
		"templates/login/js-login.html",
		"templates/login/login-form.html",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2. Render template
	data := gin.H{
		"title": "Login Page",
	}
	err = template.ExecuteTemplate(c.Writer, "view-login.html", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (ctrl *LoginController) HandleLogin(c *gin.Context) {
	// Logika penanganan login
	// ...
	// Contoh: Verifikasi kredensial, buat sesi, dll.

	// Mengarahkan pengguna setelah login berhasil
	c.Redirect(302, "/userboard")
}
