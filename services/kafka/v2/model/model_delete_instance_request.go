/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteInstanceRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
}

func (o DeleteInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}
