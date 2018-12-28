package router

import (
	"go-lyl/learn/restful/handler"

	restful "github.com/emicklei/go-restful"
)

func RegistryAPIHandler(wc *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/")
	ws.Route(ws.POST("/api/waterDevice/count").To(handler.GetWaterCount))
	ws.Route(ws.POST("/api/waterDevice/location/value").To(handler.GetWaterValue))
	ws.Route(ws.POST("/api/waterDevice/info").To(handler.GetWaterDetail))
	ws.Route(ws.POST("/api/fireDevice/count").To(handler.GetFireCount))
	ws.Route(ws.POST("/api/fireDevice/location/count").To(handler.GetFireCount))
	wc.Add(ws)
}
