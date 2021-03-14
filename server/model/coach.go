// 自动生成模板Coach
package model

// 如果含有time.Time 请自行import time包
type Coach struct {
	SysUser
}

func (Coach) TableName() string {
	return "sys_users"
}
