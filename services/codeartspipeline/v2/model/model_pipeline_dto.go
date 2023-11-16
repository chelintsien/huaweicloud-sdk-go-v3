package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineDto struct {

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 流水线描述
	Description *string `json:"description,omitempty"`

	// 是否为发布流水线
	IsPublish *bool `json:"is_publish,omitempty"`

	// 流水线源
	Sources *[]CodeSource `json:"sources,omitempty"`

	// 流水线自定义全局变量
	Variables *[]CustomVariable `json:"variables,omitempty"`

	// 流水线定时执行配置
	Schedules *[]PipelineSchedule `json:"schedules,omitempty"`

	// 流水线代码事件触发配置
	Triggers *[]PipelineTrigger `json:"triggers,omitempty"`

	// 流水线结构定义版本，新版默认为3.0
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// 流水线结构定义
	Definition *string `json:"definition,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`
}

func (o PipelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineDto struct{}"
	}

	return strings.Join([]string{"PipelineDto", string(data)}, " ")
}
