package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_EAST_2 = region.NewRegion("cn-east-2", "https://lts.cn-east-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://lts.cn-east-3.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://lts.cn-north-1.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://lts.cn-north-2.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://lts.cn-north-4.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://lts.cn-south-1.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://lts.cn-southwest-2.myhuaweicloud.com")
var AF_SOUTH_1 = region.NewRegion("af-south-1", "https://lts.af-south-1.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://lts.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://lts.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://lts.ap-southeast-3.myhuaweicloud.com")
var LA_SOUTH_2 = region.NewRegion("la-south-2", "https://lts.la-south-2.myhuaweicloud.com")
var SA_BRAZIL_1 = region.NewRegion("sa-brazil-1", "https://lts.sa-brazil-1.myhuaweicloud.com")
var LA_NORTH_2 = region.NewRegion("la-north-2", "https://lts.la-north-2.myhuaweicloud.com")
var CN_NORTH_9 = region.NewRegion("cn-north-9", "https://lts-pctest.cn-north-9.myhuaweicloud.cn")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://lts.cn-south-2.myhuaweicloud.com")
var NA_MEXICO_1 = region.NewRegion("na-mexico-1", "https://lts.na-mexico-1.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-north-1":     CN_NORTH_1,
	"cn-north-2":     CN_NORTH_2,
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"af-south-1":     AF_SOUTH_1,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"la-south-2":     LA_SOUTH_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-north-2":     LA_NORTH_2,
	"cn-north-9":     CN_NORTH_9,
	"cn-south-2":     CN_SOUTH_2,
	"na-mexico-1":    NA_MEXICO_1,
}

var provider = region.DefaultProviderChain("LTS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
