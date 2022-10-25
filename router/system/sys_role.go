package system

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/api"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/role")
	roleGroup := api.BuildGroup.SystemApiGroup.RoleApi
	{
		userRouter.GET("/pageQuery", roleGroup.PageQuery)
	}
}
