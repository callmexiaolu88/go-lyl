package main

import (
	"go-lyl/learn/restful/router"

	web "github.com/micro/go-web"
	restfulspec "github.com/use-go/go-restful-openapi"

	restful "github.com/emicklei/go-restful"
)

func main() {
	srv := web.NewService(web.Name("restful.service"),
		web.Address("127.0.0.1:6300"))
	wc := restful.NewContainer()
	router.RegistryAPIHandler(wc)

	config := restfulspec.Config{
		WebServices:    wc2.RegisteredWebServices(),
		WebServicesURL: ":8888",
		APIPath:        "/apidocs.json",
	}
	wc.Add(restfulspec.NewOpenAPIService(config))

	srv.Handle("/", wc)
	srv.Init()
	srv.Run()
}
