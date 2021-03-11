package request

import "Driving-school/model"

type CoachSearch struct {
	model.Coach
	PageInfo
}
