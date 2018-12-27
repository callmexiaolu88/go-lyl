package handler

import (
	"encoding/json"
	"fmt"
	"go-lyl/learn/restful/model"
	"net/http"

	"honeywell.com/foxconn/fire-platform-water-srv/handler"
	water "honeywell.com/foxconn/fire-platform-water-srv/proto"

	restful "github.com/emicklei/go-restful"
)

type UserResource struct {
	Users map[string]model.User
}

var users UserResource = UserResource{map[string]model.User{}}

func GetCount(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	obj := &[]handler.DeviceInfoData{
		handler.DeviceInfoData{
			Code: 200,
			Msg:  "",
			Data: []*water.DeviceInfo{
				&water.DeviceInfo{
					Name:  "wuxian",
					Value: 100,
				},
				&water.DeviceInfo{
					Name:  "yeya",
					Value: 255,
				},
				&water.DeviceInfo{
					Name:  "ast",
					Value: 320,
				},
			},
		},
	}
	body, err := json.Marshal(obj)
	if err != nil {
		body, err = json.Marshal(&handler.DeviceInfoData{
			Code: 500,
			Msg:  "marshal error",
		})
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	fmt.Println(string(body))
	rsp.Write(body)
}

func GetValue(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	obj := &[]handler.DeviceInfoData{
		handler.DeviceInfoData{
			Code: 200,
			Msg:  "",
			Data: []*water.DeviceInfo{
				&water.DeviceInfo{
					Name:  "shuiya",
					Value: 1200,
				},
				&water.DeviceInfo{
					Name:  "shuiwei",
					Value: 1500,
				},
				&water.DeviceInfo{
					Name:  "dianyuan",
					Value: 220,
				},
			},
		},
	}
	body, err := json.Marshal(obj)
	if err != nil {
		body, err = json.Marshal(&handler.DeviceInfoData{
			Code: 500,
			Msg:  "marshal error",
		})
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	rsp.Write(body)
}

func GetDetail(req *restful.Request, rsp *restful.Response) {
	fmt.Println(req.SelectedRoutePath())
	obj := &[]handler.DeviceDetailData{
		handler.DeviceDetailData{
			Code: 200,
			Msg:  "",
			Data: []*water.DeviceDetail{
				&water.DeviceDetail{
					Max:        20,
					Min:        15,
					Unit:       "cm",
					Value:      17,
					DeviceType: "shuiwei",
				},
				&water.DeviceDetail{
					Max:        35,
					Min:        2,
					Unit:       "pp",
					Value:      25,
					DeviceType: "yali",
				},
			},
		},
	}
	body, err := json.Marshal(obj)
	if err != nil {
		body, err = json.Marshal(&handler.DeviceInfoData{
			Code: 500,
			Msg:  "marshal error",
		})
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	rsp.Write(body)
}
