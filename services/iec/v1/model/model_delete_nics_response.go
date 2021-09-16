package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteNicsResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNicsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteNicsResponse struct{}"
	}

	return strings.Join([]string{"DeleteNicsResponse", string(data)}, " ")
}
