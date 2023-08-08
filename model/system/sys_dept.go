package system

import (
	"time"
)

// SysDept 部门表
type SysDept struct {
	DeptId     int64     `json:"deptId" gorm:"column:dept_id;"`         // 部门id
	ParentId   int64     `json:"parentId" gorm:"column:parent_id;"`     // 父部门id
	Ancestors  string    `json:"ancestors" gorm:"column:ancestors;"`    // 祖级列表
	DeptName   string    `json:"deptName" gorm:"column:dept_name;"`     // 部门名称
	OrderNum   int64     `json:"orderNum" gorm:"column:order_num;"`     // 显示顺序
	Leader     string    `json:"leader" gorm:"column:leader;"`          // 负责人
	Phone      string    `json:"phone" gorm:"column:phone;"`            // 联系电话
	Email      string    `json:"email" gorm:"column:email;"`            // 邮箱
	Status     string    `json:"status" gorm:"column:status;"`          // 部门状态（0正常 1停用）
	DelFlag    string    `json:"delFlag" gorm:"column:del_flag;"`       // 删除标志（0代表存在 2代表删除）
	CreateBy   string    `json:"createBy" gorm:"column:create_by;"`     // 创建者
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;"` // 创建时间
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by;"`     // 更新者
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;"` // 更新时间
}

// TableName 表名称
func (*SysDept) TableName() string {
	return "sys_dept"
}
