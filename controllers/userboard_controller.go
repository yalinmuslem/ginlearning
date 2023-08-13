package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserboardController struct{}

func (ctrl *UserboardController) Index(c *gin.Context) {

	c.HTML(200, "view-userboard.html", gin.H{
		"title": "Userboard",
	})
}
