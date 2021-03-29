package request

import "Driving-school/model"

type ExamSearch struct {
	model.Exam
	PageInfo
}
