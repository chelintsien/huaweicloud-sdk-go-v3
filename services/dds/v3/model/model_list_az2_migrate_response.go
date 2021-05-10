package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAz2MigrateResponse struct {
	// 可用区具体信息。

	AzList         *[]Az2Migrate `json:"az_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAz2MigrateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAz2MigrateResponse struct{}"
	}

	return strings.Join([]string{"ListAz2MigrateResponse", string(data)}, " ")
}
