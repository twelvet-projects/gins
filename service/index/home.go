package index

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/model/response"
)

type HomeService struct{}

// Index 首页服务
func (h *HomeService) Index(c *gin.Context) {
	response.OkWithMessage("欢迎使用Gins", c)
}
