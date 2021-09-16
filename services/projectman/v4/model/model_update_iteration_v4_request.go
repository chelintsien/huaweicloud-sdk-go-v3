package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateIterationV4Request struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 迭代id

	IterationId int32 `json:"iteration_id"`

	Body *UpdateIterationRequestV4 `json:"body,omitempty"`
}

func (o UpdateIterationV4Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateIterationV4Request struct{}"
	}

	return strings.Join([]string{"UpdateIterationV4Request", string(data)}, " ")
}
