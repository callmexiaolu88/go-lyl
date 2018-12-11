package main

import (
	"lyl/learn/micro/grpc/handler"
	mathFunction "lyl/learn/micro/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":43300")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	mathFunction.RegisterMathFunctionServer(server, handler.MathFunction{})
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
