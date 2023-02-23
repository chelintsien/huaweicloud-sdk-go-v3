package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRoleNamesResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRoleNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoleNamesResponse struct{}"
	}

	return strings.Join([]string{"ListRoleNamesResponse", string(data)}, " ")
}
