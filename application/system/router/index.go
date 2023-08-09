package router

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/application/system/controller"
)

type IndexRouter struct{}

func (s *IndexRouter) InitIndexRouter(router *gin.RouterGroup) {
	indexApi := controller.IndexApi{}
	{
		router.GET("/", indexApi.Index)
	}
}
