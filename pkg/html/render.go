package html

import (
	"github.com/gin-gonic/gin"
	"goblog/internal/providers/view"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.GlobalData(data)
	c.HTML(code, name, data)
}
