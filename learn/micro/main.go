package main

import (
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/transport/grpc"
)

func main() {
	service := micro.NewService(
		micro.Name("learn.micro"),
	)

	service.Init()
	service.Run()
}
