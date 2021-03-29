package model

import "Driving-school/global"

type Exam struct {
	global.GVA_MODEL
	UserId  uint   `json:"userId" `
	Name    string `json:"name"`
	Result  string `json:"result"`
	Message string `json:"message"`
}
