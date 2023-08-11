package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/application/system/service"
)

type IndexApi struct{}

var (
	indexService = service.IndexService{}
)

// Index
// @Summary 首页欢迎
// @Description 欢迎Json
// @Tags Index
// @Produce   application/json
// @Success 200 {object} string
// @Router / [get]
func (h *IndexApi) Index(c *gin.Context) {
	indexService.Index(c)
}
