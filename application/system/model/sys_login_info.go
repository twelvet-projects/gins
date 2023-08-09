package model

import "time"

// SysLoginInfo 系统访问记录
type SysLoginInfo struct {
	InfoId     int64     `json:"infoId" gorm:"column:info_id;"`         // 访问ID
	UserName   string    `json:"userName" gorm:"column:user_name;"`     // 用户账号
	Ipaddr     string    `json:"ipaddr" gorm:"column:ipaddr;"`          // 登录IP地址
	Status     string    `json:"status" gorm:"column:status;"`          // 登录状态（0登录成功 1登录失败 2成功退出）
	Msg        string    `json:"msg" gorm:"column:msg;"`                // 提示信息
	AccessTime time.Time `json:"accessTime" gorm:"column:access_time;"` // 访问时间
	DeptId     int64     `json:"deptId" gorm:"column:dept_id;"`         // 部门ID
}

// TableName 表名称
func (*SysLoginInfo) TableName() string {
	return "sys_login_info"
}
