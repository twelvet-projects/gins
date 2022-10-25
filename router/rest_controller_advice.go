package router

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// Advice404 处理404
func Advice404(c *gin.Context) {
	c.JSON(404, gin.H{
		"code": 404,
		"msg":  "Not Found",
	})
}

// Advice500 处理500
func Advice500(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			c.JSON(200, gin.H{
				"code":    500,
				"message": "服务器内部错误",
			})
		}
	}()
	c.Next()
}
