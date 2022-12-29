package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Ipaddresses struct {
	Ipaddress *IpaddressData `json:"ipaddress,omitempty"`
}

func (o Ipaddresses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ipaddresses struct{}"
	}

	return strings.Join([]string{"Ipaddresses", string(data)}, " ")
}
