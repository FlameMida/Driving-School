package model

import (
	"Driving-school/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID          uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username      string       `json:"userName" form:"username" gorm:"comment:用户登录名"`
	Password      string       `json:"-"  form:"password" gorm:"comment:用户登录密码"`
	NickName      string       `json:"nickName" form:"nickName" gorm:"default:系统用户;comment:用户姓名" `
	HeaderImg     string       `json:"headerImg" form:"headerImg" gorm:"default:;comment:用户头像"`
	Authority     SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId   string       `json:"authorityId" form:"authorityId" gorm:"default:888;comment:用户角色ID"`
	HasServerInfo bool         `json:"has_server_info" gorm:"-" `
	CoachId       uint         `json:"coachID" gorm:"coachID;comment:教练id;default:0"`
	Phone         string       `json:"phone" gorm:"phone;comment:手机号码"`
}
