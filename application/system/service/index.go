package service

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/framework/model/response"
)

type IndexService struct{}

// Index 首页服务
func (h *IndexService) Index(c *gin.Context) {
	response.OkWithMessage("欢迎使用Gins", c)
}
