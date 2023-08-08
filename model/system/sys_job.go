package system

import "time"

// SysJob 定时任务调度表
type SysJob struct {
	JobId          int64     `json:"jobId" gorm:"column:job_id;"`                   // 任务ID
	JobName        string    `json:"jobName" gorm:"column:job_name;"`               // 任务名称
	JobGroup       string    `json:"jobGroup" gorm:"column:job_group;"`             // 任务组名
	InvokeTarget   string    `json:"invokeTarget" gorm:"column:invoke_target;"`     // 调用目标字符串
	CronExpression string    `json:"cronExpression" gorm:"column:cron_expression;"` // cron执行表达式
	MisfirePolicy  int8      `json:"misfirePolicy" gorm:"column:misfire_policy;"`   // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     int8      `json:"concurrent" gorm:"column:concurrent;"`          // 是否并发执行（1允许 0禁止）
	Status         string    `json:"status" gorm:"column:status;"`                  // 状态（0正常 1暂停）
	CreateBy       string    `json:"createBy" gorm:"column:create_by;"`             // 创建者
	CreateTime     time.Time `json:"createTime" gorm:"column:create_time;"`         // 创建时间
	UpdateBy       string    `json:"updateBy" gorm:"column:update_by;"`             // 更新者
	UpdateTime     time.Time `json:"updateTime" gorm:"column:update_time;"`         // 更新时间
	Remark         string    `json:"remark" gorm:"column:remark;"`                  // 备注信息
}

// TableName 表名称
func (*SysJob) TableName() string {
	return "sys_job"
}
