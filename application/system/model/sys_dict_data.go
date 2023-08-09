package model

import "time"

// SysDictData 字典数据表
type SysDictData struct {
	DictCode   int64     `json:"dictCode" gorm:"column:dict_code;"`     // 字典编码
	DictSort   int64     `json:"dictSort" gorm:"column:dict_sort;"`     // 字典排序
	DictLabel  string    `json:"dictLabel" gorm:"column:dict_label;"`   // 字典标签
	DictValue  string    `json:"dictValue" gorm:"column:dict_value;"`   // 字典键值
	DictType   string    `json:"dictType" gorm:"column:dict_type;"`     // 字典类型
	CssClass   string    `json:"cssClass" gorm:"column:css_class;"`     // 样式属性（其他样式扩展）
	ListClass  string    `json:"listClass" gorm:"column:list_class;"`   // 表格回显样式
	IsDefault  string    `json:"isDefault" gorm:"column:is_default;"`   // 是否默认（Y是 N否）
	Status     int8      `json:"status" gorm:"column:status;"`          // 状态（0正常 1停用）
	CreateBy   string    `json:"createBy" gorm:"column:create_by;"`     // 创建者
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;"` // 创建时间
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by;"`     // 更新者
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;"` // 更新时间
	Remark     string    `json:"remark" gorm:"column:remark;"`          // 备注
}

// TableName 表名称
func (*SysDictData) TableName() string {
	return "sys_dict_data"
}
