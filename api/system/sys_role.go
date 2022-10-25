package system

import "github.com/gin-gonic/gin"

type RoleApi struct{}

// PageQuery 查询角色列表
func (r *RoleApi) PageQuery(c *gin.Context) {
	c.JSON(200, "成功")
}
