package model

import "time"

// SysOperationLog 操作日志记录
type SysOperationLog struct {
	OperId        int64     `json:"operId" gorm:"column:oper_id;"`               // 日志主键
	Service       string    `json:"service" gorm:"column:service;"`              // 模块标题
	BusinessType  int64     `json:"businessType" gorm:"column:business_type;"`   // 业务类型（0其它 1新增 2修改 3删除）
	Method        string    `json:"method" gorm:"column:method;"`                // 方法名称
	RequestMethod string    `json:"requestMethod" gorm:"column:request_method;"` // 请求方式
	OperatorType  int8      `json:"operatorType" gorm:"column:operator_type;"`   // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string    `json:"operName" gorm:"column:oper_name;"`           // 操作人员
	DeptName      string    `json:"deptName" gorm:"column:dept_name;"`           // 部门名称
	OperUrl       string    `json:"operUrl" gorm:"column:oper_url;"`             // 请求URL
	OperIp        string    `json:"operIp" gorm:"column:oper_ip;"`               // 主机地址
	OperLocation  string    `json:"operLocation" gorm:"column:oper_location;"`   // 操作地点
	OperParam     string    `json:"operParam" gorm:"column:oper_param;"`         // 请求参数
	JsonResult    string    `json:"jsonResult" gorm:"column:json_result;"`       // 返回参数
	Status        int8      `json:"status" gorm:"column:status;"`                // 操作状态（0正常 1异常）
	ErrorMsg      string    `json:"errorMsg" gorm:"column:error_msg;"`           // 错误消息
	OperTime      time.Time `json:"operTime" gorm:"column:oper_time;"`           // 操作时间
	DeptId        int64     `json:"deptId" gorm:"column:dept_id;"`               // 部门ID
}

// TableName 表名称
func (*SysOperationLog) TableName() string {
	return "sys_operation_log"
}
