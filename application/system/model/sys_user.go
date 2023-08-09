package model

import "time"

// SysUser 用户信息
type SysUser struct {
	UserId      int64     `json:"userId" gorm:"column:user_id;primarykey;AUTO_INCREMENT;type:bigint;size:20;comment:用户ID;"` // 用户ID
	DeptId      int64     `json:"deptId" gorm:"column:deptid;type:bigint;size:20;comment:部门ID;"`                            // 部门ID
	Username    string    `json:"username" gorm:"column:username;not null;size:30;comment:用户账号;"`                           // 用户账号
	NickName    string    `json:"nickName" gorm:"column:nickname;not null;size:30;comment:用户昵称;"`                           // 用户昵称
	UserType    string    `json:"userType" gorm:"column:usertype;default:00;size:2;comment:用户类型（00系统用户）;"`                  // 用户类型（00系统用户）
	Email       string    `json:"email" gorm:"column:email;default:'';size:50;comment:用户邮箱;"`                               // 用户邮箱
	Phonenumber string    `json:"phonenumber" gorm:"column:phonenumber;default:'';size:11;comment:手机号码;"`                   // 手机号码
	Sex         int8      `json:"sex" gorm:"column:sex;size:1;comment:用户性别（0男 1女 2未知）;"`                                    // 用户性别（0男 1女 2未知）
	Avatar      string    `json:"avatar" gorm:"column:avatar;default:'';size:100;comment:头像地址;"`                            // 头像地址
	Password    string    `json:"password" gorm:"column:password;default:'';size:100;comment:密码;"`                          // 密码
	Status      int8      `json:"status" gorm:"column:status;size:1;comment:帐号状态（0正常 1停用）;"`                                // 帐号状态（0正常 1停用）
	DelFlag     string    `json:"delFlag" gorm:"column:delflag;default:0;size:1;comment:删除标志（0代表存在 2代表删除）;"`                // 删除标志（0代表存在 2代表删除）
	LoginIp     string    `json:"loginIp" gorm:"column:loginip;default:'';size:50;comment:最后登陆IP;"`                         // 最后登陆IP
	LoginDate   time.Time `json:"loginDate" gorm:"column:logindate;comment:最后登陆时间;"`                                        // 最后登陆时间
	CreateBy    string    `json:"createBy" gorm:"column:createby;default:'';size:64;comment:创建者;"`                          // 创建者
	CreateTime  time.Time `json:"createTime" gorm:"column:createtime;comment:创建时间;"`                                        // 创建时间
	UpdateBy    string    `json:"updateBy" gorm:"column:updateby;default:'';size:64;comment:更新者;"`                          // 更新者
	UpdateTime  time.Time `json:"updateTime" gorm:"column:updatetime;comment:更新时间;"`                                        // 更新时间
	Remark      string    `json:"remark" gorm:"column:remark;size:500;comment:备注;"`                                         // 备注
}

// TableName 表名称
func (*SysUser) TableName() string {
	return "sys_user"
}
