package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
}

// StartServer 启动服务
func StartServer(port string, router *gin.Engine) Server {
	return &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
