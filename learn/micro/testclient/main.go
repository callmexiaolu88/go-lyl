package main

import (
	"context"
	"encoding/json"
	"fmt"
	mathFunction "lyl/learn/micro/proto"

	"github.com/micro/go-micro/client"
)

type testStr struct {
	Name    string `json:"ID"`
	id      string `json:"ID"`
	Address string `json:"-"`
	phone   string `json:"-"`
}

func main() {
	/*service := micro.NewService(micro.Name("learn.client"))
	service.Init()
	cli := mathFunction.NewMathFunctionService("learn.micro", service.Client())*/
	cli := mathFunction.NewMathFunctionService("learn.micro", client.DefaultClient)
	result, err := cli.Add(context.Background(), &mathFunction.AddRequest{
		P1: 10,
		P2: 20,
		P3: 30,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Result)
	fmt.Println("==============")
	fmt.Println(result)
	r1, err := cli.Subtraction(context.Background(), &mathFunction.SubtractionRequest{
		P1: 1000,
		P2: 200,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r1.Result)
	fmt.Println("==============")
	fmt.Println(r1)
	fmt.Println("==============")
	st := testStr{
		Name:    "lyl",
		id:      "12",
		Address: "shanghai",
		phone:   "136971255",
	}
	fmt.Println(string(func() []byte { v, _ := json.Marshal(st); return v }()))
}
