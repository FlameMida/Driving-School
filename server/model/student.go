package model

type Student struct {
	SysUser
}

func (Student) TableName() string {
	return "sys_users"
}
