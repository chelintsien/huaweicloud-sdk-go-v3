package model

import (
	"encoding/json"

	"strings"
)

type SqlserverGrantRequest struct {
	// 数据库名称。

	DbName string `json:"db_name"`
	// 每个元素都是与数据库相关联的帐号。单次请求最多支持50个元素。

	Users []SqlserverUserWithPrivilege `json:"users"`
}

func (o SqlserverGrantRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SqlserverGrantRequest struct{}"
	}

	return strings.Join([]string{"SqlserverGrantRequest", string(data)}, " ")
}
