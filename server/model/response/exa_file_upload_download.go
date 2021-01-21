package response

import "Driving-school/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
