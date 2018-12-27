package main

import (
	"encoding/json"
	"fmt"
	"go-lyl/learn/restful/router"

	web "github.com/micro/go-web"
	"honeywell.com/foxconn/fire-platform-water-srv/handler"
	water "honeywell.com/foxconn/fire-platform-water-srv/proto"

	restful "github.com/emicklei/go-restful"
)

func main() {

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
	body, _ := json.Marshal(obj)
	fmt.Println(string(body))

	srv := web.NewService(web.Name("restful.service"),
		web.Address("127.0.0.1:6200"))
	wc := restful.NewContainer()
	router.RegistryAPIHandler(wc)

	srv.Handle("/", wc)
	srv.Init()
	srv.Run()
}
