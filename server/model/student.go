// 自动生成模板Student
package model

// 如果含有time.Time 请自行import time包
type Student struct {
	SysUser
}

func (Student) TableName() string {
	return "sys_users"
}
