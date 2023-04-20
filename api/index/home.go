package index

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/service/index"
)

type HomeApi struct{}

var (
	homeService = index.HomeService{}
)

// Index 首页Api
func (h *HomeApi) Index(c *gin.Context) {
	homeService.Index(c)
}
