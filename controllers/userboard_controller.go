package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserboardController struct{}

func (ctrl *UserboardController) Index(c *gin.Context) {

	// 1. Load template
	template, err := template.ParseFiles(
		"templates/view-userboard.html",
		"templates/userboard/css-userboard.html",
		"templates/userboard/js-userboard.html",
		"templates/userboard/userboard.html",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2. Render template
	data := gin.H{
		"title": "Userboard Page",
	}
	err = template.ExecuteTemplate(c.Writer, "view-userboard.html", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
