package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteVpcTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcTagResponse", string(data)}, " ")
}
