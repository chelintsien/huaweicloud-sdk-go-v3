package model

import (
	"encoding/json"

	"strings"
)

//
type ResourceTag struct {
	// 功能说明：标签键 约束：同一资源的key值不能重复。

	Key string `json:"key"`
	// 功能说明：标签值

	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
