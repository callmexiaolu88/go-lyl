package main

import (
	"context"
	"fmt"
	mathFunction "lyl/learn/micro/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:43300", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := mathFunction.NewMathFunctionClient(conn)
	result, err := client.Add(context.Background(), &mathFunction.AddRequest{
		P1: 10,
		P2: 20,
		P3: 30,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Result)

	result1, err := client.Subtraction(context.Background(), &mathFunction.SubtractionRequest{
		P1: 10,
		P2: 20,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result1.Result)
}
