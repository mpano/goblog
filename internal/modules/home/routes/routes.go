package routes

import (
	"github.com/gin-gonic/gin"
	homeCtrl "goblog/internal/modules/home/controllers"
	"goblog/pkg/html"
	"net/http"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/article", gin.H{
			"title": "Article Page",
		})
	})
	router.GET("/login", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/login", gin.H{
			"title": "Login Page",
		})
	})
	router.GET("/register", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/register", gin.H{
			"title": "Register Page",
		})
	})
	router.GET("/article/create", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/create-article", gin.H{
			"title": "Create Article Page",
		})
	})
}
