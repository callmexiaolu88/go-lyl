package handler

import (
	"encoding/json"
	"fmt"
	"go-lyl/learn/restful/model"
	"io/ioutil"

	"honeywell.com/foxconn/fire-platform-common-pkg/protocol/util"

	restful "github.com/emicklei/go-restful"
)

func GetCount(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	bytes, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		return
	}
	result := model.DeviceCountResponseInfo{
		Code: 200,
		Msg:  "Query success",
	}
	conds := &[]model.DeviceCountCondition{}
	err = json.Unmarshal(bytes, conds)
	if err != nil {
		return
	}
	result.Data = make([]model.DeviceCountConditionAndResult, len(*conds))
	for k, c := range *conds {
		result.Data[k].Search = c
		result.Data[k].Values = []model.DeviceCountResult{
			model.DeviceCountResult{
				Name:  util.RandString(6),
				Value: util.RandCount(1000),
			},
			model.DeviceCountResult{
				Name:  util.RandString(6),
				Value: util.RandCount(1000),
			},
			model.DeviceCountResult{
				Name:  util.RandString(6),
				Value: util.RandCount(1000),
			},
		}
	}
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	rsp.Write(body)
}

func GetWaterValue(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	bytes, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		return
	}
	result := model.DeviceWaterValueResponseInfo{
		Code: 200,
		Msg:  "Query success",
	}
	conds := &[]model.DeviceValueCondition{}
	err = json.Unmarshal(bytes, conds)
	if err != nil {
		return
	}
	result.Data = make([]model.DeviceWaterValueConditionAndResult, len(*conds))
	for k, c := range *conds {
		result.Data[k].Search = c
		result.Data[k].Values = []model.DeviceWaterValueResult{
			model.DeviceWaterValueResult{
				DeviceID:    util.RandString(6),
				DeviceLabel: util.RandString(6),
				Value:       int32(util.RandCount(100)),
				DeviceType:  util.RandString(6),
				MapLocation: getmp(),
			},
			model.DeviceWaterValueResult{
				DeviceID:    util.RandString(6),
				DeviceLabel: util.RandString(6),
				Value:       int32(util.RandCount(100)),
				DeviceType:  util.RandString(6),
				MapLocation: getmp(),
			},
			model.DeviceWaterValueResult{
				DeviceID:    util.RandString(6),
				DeviceLabel: util.RandString(6),
				Value:       int32(util.RandCount(100)),
				DeviceType:  util.RandString(6),
				MapLocation: getmp(),
			},
		}
	}
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	rsp.Write(body)
}

func GetFireValue(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	bytes, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		return
	}
	result := model.DeviceFireValueResponseInfo{
		Code: 200,
		Msg:  "Query success",
	}
	conds := &[]model.DeviceValueCondition{}
	err = json.Unmarshal(bytes, conds)
	if err != nil {
		return
	}
	result.Data = make([]model.DeviceFireValueConditionAndResult, len(*conds))
	for k, c := range *conds {
		result.Data[k].Search = c
		result.Data[k].Values = []model.DeviceFireValueResult{
			model.DeviceFireValueResult{
				DeviceID:     util.RandString(6),
				DeviceLabel:  util.RandString(6),
				DeviceType:   util.RandString(6),
				MapLocation:  getmp(),
				DeviceStatus: []string{util.RandString(6), util.RandString(6), util.RandString(6)},
			},
			model.DeviceFireValueResult{
				DeviceID:     util.RandString(6),
				DeviceLabel:  util.RandString(6),
				DeviceType:   util.RandString(6),
				MapLocation:  getmp(),
				DeviceStatus: []string{util.RandString(6), util.RandString(6), util.RandString(6)},
			},
			model.DeviceFireValueResult{
				DeviceID:     util.RandString(6),
				DeviceLabel:  util.RandString(6),
				DeviceType:   util.RandString(6),
				MapLocation:  getmp(),
				DeviceStatus: []string{util.RandString(6), util.RandString(6), util.RandString(6)},
			},
		}
	}
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	rsp.Write(body)
}

func GetDetail(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	bytes, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		return
	}
	result := model.DeviceDetailResponseInfo{
		Code: 200,
		Msg:  "Query success",
	}
	conds := &[]model.DeviceDetailCondition{}
	err = json.Unmarshal(bytes, conds)
	if err != nil {
		return
	}
	result.Data = make([]model.DeviceDetailConditionAndResult, len(*conds))
	for k, c := range *conds {
		result.Data[k].Search = c
		result.Data[k].Values = []model.DeviceDetailResult{
			model.DeviceDetailResult{
				Max:        int32(util.RandCount(100)),
				Min:        int32(util.RandCount(6)),
				Value:      int32(util.RandCount(100)),
				DeviceType: util.RandString(6),
				Unit:       util.RandString(3),
			},
			model.DeviceDetailResult{
				Max:        int32(util.RandCount(100)),
				Min:        int32(util.RandCount(6)),
				Value:      int32(util.RandCount(100)),
				DeviceType: util.RandString(6),
				Unit:       util.RandString(3),
			},
		}
	}
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	rsp.Write(body)
}

func getmp() *model.LocationInfo {
	m := &model.LocationInfo{}
	m.F1 = []float32{float32(util.RandCount(1000)), float32(util.RandCount(1000)), float32(util.RandCount(1000))}
	m.S1 = []float32{float32(util.RandCount(1000)), float32(util.RandCount(1000)), float32(util.RandCount(1000))}
	m.V1 = []float32{float32(util.RandCount(1000)), float32(util.RandCount(1000)), float32(util.RandCount(1000))}
	return m
}
