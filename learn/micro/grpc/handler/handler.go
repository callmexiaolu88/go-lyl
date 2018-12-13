package handler

import (
	"context"
	"log"
	mathFunction "lyl/learn/micro/proto"
)

//MathFunction is a struct
type MathFunction struct{}

//Add return
func (MathFunction) Add(context context.Context, req *mathFunction.AddRequest) (*mathFunction.AddResponse, error) {
	log.Printf("%s\n", context)
	log.Printf("%s\n", req)
	return &mathFunction.AddResponse{
		Result: req.P1 + req.P2 + req.P3,
	}, nil
}

//Subtraction return
func (MathFunction) Subtraction(context context.Context, req *mathFunction.SubtractionRequest) (*mathFunction.SubtractionResponse, error) {
	log.Printf("%s\n", context)
	log.Printf("%s\n", req)
	return &mathFunction.SubtractionResponse{
		Result: req.P1 - req.P2,
	}, nil
}
