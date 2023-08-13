// controllers/login_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

// Di dalam file login_controller.go
func (ctrl *LoginController) ShowLoginForm(c *gin.Context) {
	c.HTML(200, "view-login.html", gin.H{
		"title": "Login Form",
	})
}

func (ctrl *LoginController) HandleLogin(c *gin.Context) {
	// Logika penanganan login
	// ...
	// Contoh: Verifikasi kredensial, buat sesi, dll.

	// Mengarahkan pengguna setelah login berhasil
	c.Redirect(302, "/userboard")
}
