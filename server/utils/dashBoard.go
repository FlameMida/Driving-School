package utils

import "time"

type DashBoard struct {
	TimeInfo      string `json:"time_info"`
	TimeMsg       string `json:"time_msg"`
	ActiveUsers   string `json:"active_users"`
	TotalUsers    int    `json:"total_users"`
	TotalStudents int    `json:"total_students"`
}

func InitTimeInfo() (info string, msg string) {
	Hour := time.Now().Hour()
	if Hour >= 23 && Hour < 6 {
		info = "夜深了"
		msg = "请您注意休息~祝您晚安！"
	} else if Hour >= 6 && Hour < 12 {
		info = "早上好"
		msg = "请开始您一天的工作吧！"
	} else if Hour >= 12 && Hour < 14 {
		info = "中午好"
		msg = "祝您有一个好的胃口！"
	} else if Hour >= 14 && Hour < 18 {
		info = "下午好"
		msg = "祝愿您有美好的一天~！"
	} else {
		info = "晚上好"
		msg = "祝您今夜好梦~！"
	}
	return info, msg
}
