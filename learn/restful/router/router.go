package router

import (
	"github.com/use-go/go-restful-openapi"
	"lyl/learn/restful/handler"

	restful "github.com/emicklei/go-restful"
)

func RegistryAPIHandler(wc *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/user")
	ws.Route(ws.GET("/build/{id}").
		Doc("this is get user function").
		Metadata(restfulspec.KeyOpenAPITags,[]string{"get","id"}).
		To(handler.GetUser))
	wc.Add(ws)
}
