package model

import "time"

// SysMenu 菜单权限表
type SysMenu struct {
	MenuId     int64     `json:"menuId" gorm:"column:menu_id;"`         // 菜单ID
	MenuName   string    `json:"menuName" gorm:"column:menu_name;"`     // 菜单名称
	ParentId   int64     `json:"parentId" gorm:"column:parent_id;"`     // 父菜单ID
	OrderNum   int64     `json:"orderNum" gorm:"column:order_num;"`     // 显示顺序
	Path       string    `json:"path" gorm:"column:path;"`              // 路由地址
	Component  string    `json:"component" gorm:"column:component;"`    // 组件路径
	IsFrame    string    `json:"isFrame" gorm:"column:is_frame;"`       // 是否为外链（0是 1否）
	MenuType   string    `json:"menuType" gorm:"column:menu_type;"`     // 菜单类型（M目录 C菜单 F按钮）
	Visible    string    `json:"visible" gorm:"column:visible;"`        // 菜单状态（0显示 1隐藏）
	Status     string    `json:"status" gorm:"column:status;"`          // 菜单状态（0正常 1停用）
	Perms      string    `json:"perms" gorm:"column:perms;"`            // 权限标识
	Icon       string    `json:"icon" gorm:"column:icon;"`              // 菜单图标
	CreateBy   string    `json:"createBy" gorm:"column:create_by;"`     // 创建者
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;"` // 创建时间
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by;"`     // 更新者
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;"` // 更新时间
	Remark     string    `json:"remark" gorm:"column:remark;"`          // 备注
}

// TableName 表名称
func (*SysMenu) TableName() string {
	return "sys_menu"
}
