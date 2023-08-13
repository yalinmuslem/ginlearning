package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExploreJobsController struct{}

func (ctrl *ExploreJobsController) Index(c *gin.Context) {
	template, err := template.ParseFiles(
		"templates/view-userboard.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/menu.html",
		"templates/explore-jobs/css.html",
		"templates/explore-jobs/js.html",
		"templates/explore-jobs/explore-jobs.html",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := gin.H{
		"title":      "Explore Jobs",
		"activeMenu": "explore-jobs",
	}
	err = template.ExecuteTemplate(c.Writer, "view-userboard.html", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
