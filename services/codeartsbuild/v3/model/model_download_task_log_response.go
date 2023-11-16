package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadTaskLogResponse Response Object
type DownloadTaskLogResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadTaskLogResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadTaskLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadTaskLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadTaskLogResponse", string(data)}, " ")
}
