package model

type Coach struct {
	SysUser
}

func (Coach) TableName() string {
	return "sys_users"
}
