package model

import "time"

// SysPost 岗位信息表
type SysPost struct {
	PostId     int64     `json:"postId" gorm:"column:post_id;"`         // 岗位ID
	PostCode   string    `json:"postCode" gorm:"column:post_code;"`     // 岗位编码
	PostName   string    `json:"postName" gorm:"column:post_name;"`     // 岗位名称
	PostSort   int64     `json:"postSort" gorm:"column:post_sort;"`     // 显示顺序
	Status     string    `json:"status" gorm:"column:status;"`          // 状态（0正常 1停用）
	CreateBy   string    `json:"createBy" gorm:"column:create_by;"`     // 创建者
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;"` // 创建时间
	UpdateBy   string    `json:"updateBy" gorm:"column:update_by;"`     // 更新者
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;"` // 更新时间
	Remark     string    `json:"remark" gorm:"column:remark;"`          // 备注
}

// TableName 表名称
func (*SysPost) TableName() string {
	return "sys_post"
}
