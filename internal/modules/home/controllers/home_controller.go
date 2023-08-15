package controllers

import (
	"github.com/gin-gonic/gin"
	"goblog/pkg/html"
	"net/http"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (controller Controller) Index(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home Page",
	})
}
