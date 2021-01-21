package response

import "Driving-school/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
