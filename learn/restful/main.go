package main

import (
	"go-lyl/learn/restful/router"

	web "github.com/micro/go-web"

	restful "github.com/emicklei/go-restful"
)

func main() {
	srv := web.NewService(web.Name("restful.service"),
		web.Address("127.0.0.1:6200"))
	wc := restful.NewContainer()
	router.RegistryAPIHandler(wc)

	srv.Handle("/", wc)
	srv.Init()
	srv.Run()
}
