package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "goblog/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}
