package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEdgeCloudResponse struct {
	// 边缘业务数量。

	Count *int32 `json:"count,omitempty"`
	// 边缘业务列表。

	Edgeclouds     *[]EdgeCloud `json:"edgeclouds,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEdgeCloudResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEdgeCloudResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeCloudResponse", string(data)}, " ")
}
