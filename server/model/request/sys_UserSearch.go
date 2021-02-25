package request

import "Driving-school/model"

type UserSearch struct {
	model.SysUser
	PageInfo
}
