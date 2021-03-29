package request

import "Driving-school/model"
import uuid "github.com/satori/go.uuid"

type StudentSearch struct {
	model.Student
	PageInfo
}

// Modify  user's coach structure
type SetUserCoach struct {
	UUID    uuid.UUID `json:"uuid"`
	CoachId uint      `json:"coachId"`
}
