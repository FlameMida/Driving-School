package request

import "Driving-school/model"

type StudentSearch struct {
	model.Student
	PageInfo
}
