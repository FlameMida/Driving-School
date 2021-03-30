package model

import (
	"Driving-school/global"
	"time"
)

type Exam struct {
	global.GVA_MODEL
	UserId  uint      `json:"userId" form:"userId"`
	Name    string    `json:"name" form:"name"`
	Result  string    `json:"result"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
