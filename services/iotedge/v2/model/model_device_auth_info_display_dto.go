package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点认证方式信息。
type DeviceAuthInfoDisplayDto struct {

	// 边缘节点认证方式。
	AuthType string `json:"auth_type"`

	LocalPath *CertificateLocalPathDto `json:"local_path,omitempty"`
}

func (o DeviceAuthInfoDisplayDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceAuthInfoDisplayDto struct{}"
	}

	return strings.Join([]string{"DeviceAuthInfoDisplayDto", string(data)}, " ")
}
