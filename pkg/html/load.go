package html

import "github.com/gin-gonic/gin"

func LoadHTLM(router *gin.Engine) {
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
