package model

import "time"

// SysJobLog 定时任务调度日志表
type SysJobLog struct {
	JobLogId      int64     `json:"jobLogId" gorm:"column:job_log_id;"`          // 任务日志ID
	JobName       string    `json:"jobName" gorm:"column:job_name;"`             // 任务名称
	JobGroup      string    `json:"jobGroup" gorm:"column:job_group;"`           // 任务组名
	InvokeTarget  string    `json:"invokeTarget" gorm:"column:invoke_target;"`   // 调用目标字符串
	JobMessage    string    `json:"jobMessage" gorm:"column:job_message;"`       // 日志信息
	Status        string    `json:"status" gorm:"column:status;"`                // 执行状态（0正常 1失败）
	ExceptionInfo string    `json:"exceptionInfo" gorm:"column:exception_info;"` // 异常信息
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time;"`       // 创建时间
}

// TableName 表名称
func (*SysJobLog) TableName() string {
	return "sys_job_log"
}
