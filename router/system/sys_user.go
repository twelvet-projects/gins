package system

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/api"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userApi := api.BuildGroup.SystemApiGroup.UserApi
	{
		userRouter.GET("/pageQuery", userApi.PageQuery)
	}
}
