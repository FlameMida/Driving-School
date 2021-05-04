package request

import "Driving-school/model"

type ExamSearch struct {
	model.SysExam
	PageInfo
}
