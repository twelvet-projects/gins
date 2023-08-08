package system

import "time"

// SysRole 角色信息表
type SysRole struct {
	RoleId     int64     `json:"roleId" gorm:"column:role_id;"`         // 角色ID
	RoleName   string    `json:"roleName" gorm:"column:role_name;"`     // 角色名称
	RoleKey    string    `json:"roleKey" gorm:"column:role_key;"`       // 角色权限字符串
	RoleSort   int64     `json:"roleSort" gorm:"column:role_sort;"`     // 显示顺序
	DataScope  string    `json:"dataScope" gorm:"column:data_scope;"`   // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     int8      `json:"status" gorm:"column:status;"`          // 角色状态（0正常 1停用）
	DelFlag    string    `json:"delFlag" gorm:"column:del_flag;"`       // 删除标志（0代表存在 2代表删除）
	CreateBy   string    `json:"createBy" gorm:"column:create_by;"`     // 创建者
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;"` // 创建时间
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by;"`     // 更新者
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;"` // 更新时间
	Remark     string    `json:"remark" gorm:"column:remark;"`          // 备注
}

// TableName 表名称
func (*SysRole) TableName() string {
	return "sys_role"
}
