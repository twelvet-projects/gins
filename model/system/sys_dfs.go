package system

import "time"

// SysDfs 分布式文件表
type SysDfs struct {
	FileId           int64     `json:"fileId" gorm:"column:file_id;"`                      // 主键ID
	FileName         string    `json:"fileName" gorm:"column:file_name;"`                  // 文件名称
	OriginalFileName string    `json:"originalFileName" gorm:"column:original_file_name;"` // 文件原名称
	SpaceName        string    `json:"spaceName" gorm:"column:space_name;"`                // 文件分组
	Path             string    `json:"path" gorm:"column:path;"`                           // 文件路径
	Size             int64     `json:"size" gorm:"column:size;"`                           // 文件大小
	Type             string    `json:"type" gorm:"column:type;"`                           // 文件类型
	CreateTime       time.Time `json:"createTime" gorm:"column:create_time;"`              // 创建时间
}

// TableName 表名称
func (*SysDfs) TableName() string {
	return "sys_dfs"
}
