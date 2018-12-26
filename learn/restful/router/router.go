package router

import (
	"go-lyl/learn/restful/handler"

	restful "github.com/emicklei/go-restful"
)

func RegistryAPIHandler(wc *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/")
	ws.Route(ws.GET("/api/waterDevice/count").To(handler.GetCount))
	ws.Route(ws.GET("/api/waterDevice/location/value").To(handler.GetValue))
	ws.Route(ws.GET("/api/waterDevice/info").To(handler.GetDetail))
	wc.Add(ws)
}
