package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/application/system/service"
)

type IndexApi struct{}

var (
	indexService = service.IndexService{}
)

// Index 首页Api
func (h *IndexApi) Index(c *gin.Context) {
	indexService.Index(c)
}
