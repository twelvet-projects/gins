package index

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/api/index"
)

type HomeRouter struct{}

func (s *HomeRouter) InitRouter(router *gin.RouterGroup) {
	indexApi := index.HomeApi{}
	{
		router.GET("/", indexApi.Index)
	}
}
