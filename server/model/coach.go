// 自动生成模板Coach
package model

import (
	"Driving-school/global"
)

// 如果含有time.Time 请自行import time包
type Coach struct {
	global.GVA_MODEL
	AuthorityId   string `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID;type:varchar(90);size:90;default:100"`
	HasServerInfo *bool  `json:"hasServerInfo" form:"hasServerInfo" gorm:"column:has_server_info;comment:;type:tinyint;"`
	HeaderImg     string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;type:varchar(191);size:191;"`
	NickName      string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;type:varchar(191);size:191;"`
	Password      string `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;type:varchar(191);size:191;"`
	Username      string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;type:varchar(191);size:191;"`
	Uuid          string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;type:varchar(191);size:191;"`
}

func (Coach) TableName() string {
	return "sys_users"
}
