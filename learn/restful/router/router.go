package router

import (
	"go-lyl/learn/restful/handler"

	restful "github.com/emicklei/go-restful"
)

func RegistryAPIHandler(wc *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/fireiot")
	ws.Route(ws.POST("/api/waterDevice/count").To(handler.GetCount))
	ws.Route(ws.POST("/api/waterDevice/location/value").To(handler.GetWaterValue))
	ws.Route(ws.POST("/api/waterDevice/info").To(handler.GetDetail))
	ws.Route(ws.POST("/api/fireDevice/count").To(handler.GetCount))
	ws.Route(ws.POST("/api/fireDevice/location/value").To(handler.GetFireValue))
	ws.Route(ws.POST("/api/fireDevice/info").To(handler.GetDetail))
	wc.Add(ws)
}
