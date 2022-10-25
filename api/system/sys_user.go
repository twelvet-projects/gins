package system

import (
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// PageQuery 查询用户列表
func (u *UserApi) PageQuery(c *gin.Context) {
	c.JSON(200, "成功")
}
