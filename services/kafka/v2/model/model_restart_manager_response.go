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

// Response Object
type RestartManagerResponse struct {
}

func (o RestartManagerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestartManagerResponse", string(data)}, " ")
}
