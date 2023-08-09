package model

import "time"

// SysDictType 字典类型表
type SysDictType struct {
	DictId     int64     `json:"dictId" gorm:"column:dict_id;"`         // 字典主键
	DictName   string    `json:"dictName" gorm:"column:dict_name;"`     // 字典名称
	DictType   string    `json:"dictType" gorm:"column:dict_type;"`     // 字典类型
	Status     int8      `json:"status" gorm:"column:status;"`          // 状态（0正常 1停用）
	CreateBy   string    `json:"createBy" gorm:"column:create_by;"`     // 创建者
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;"` // 创建时间
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by;"`     // 更新者
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;"` // 更新时间
	Remark     string    `json:"remark" gorm:"column:remark;"`          // 备注
}

// TableName 表名称
func (*SysDictType) TableName() string {
	return "sys_dict_type"
}
