package main

import (
	"lyl/learn/restful/router"

	web "github.com/micro/go-web"
	restfulspec "github.com/use-go/go-restful-openapi"

	restful "github.com/emicklei/go-restful"
)

func main() {
	srv := web.NewService(web.Name("restful.service"),
		web.Address(":8345"))
	wc := restful.NewContainer()
	wc2 := restful.NewContainer()
	router.RegistryAPIHandler(wc)
	router.RegistryAPIHandler(wc2)

	config := restfulspec.Config{
		WebServices:    wc2.RegisteredWebServices(),
		WebServicesURL: ":8888",
		APIPath:        "/apidocs.json",
	}
	wc.Add(restfulspec.NewOpenAPIService(config))
	wc2.Add(restfulspec.NewOpenAPIService(config))

	srv.Handle("/", wc)
	srv.Handle("/test", wc2)
	srv.Init()
	srv.Run()
}
