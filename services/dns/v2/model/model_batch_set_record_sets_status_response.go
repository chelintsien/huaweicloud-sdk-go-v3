package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchSetRecordSetsStatusResponse struct {

	// 待删除zone类型，当前仅支持public或private。
	Recordsets *[]RecordsetData `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchSetRecordSetsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusResponse", string(data)}, " ")
}
