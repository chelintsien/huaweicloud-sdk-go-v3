package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRuleAclUsingPutRequest struct {

	// 规则id
	AclRuleId string `json:"acl_rule_id"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	Body *OrderRuleAclDto `json:"body,omitempty"`
}

func (o ListRuleAclUsingPutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclUsingPutRequest struct{}"
	}

	return strings.Join([]string{"ListRuleAclUsingPutRequest", string(data)}, " ")
}
