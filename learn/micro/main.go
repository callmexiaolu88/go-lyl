package main

import (
	"lyl/learn/micro/handler"
	mathFunction "lyl/learn/micro/proto"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/transport/grpc"
)

func main() {
	service := micro.NewService(
		micro.Name("learn.micro"),
	)
	mathFunction.RegisterMathFunctionHandler(service.Server(), new(handler.MathFunction))
	service.Init()
	service.Run()
}
