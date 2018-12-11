package handler

import (
	"context"
	"log"
	mathFunction "lyl/learn/micro/proto"
)

type MathFunction struct{}

func (MathFunction) Add(context context.Context, req *mathFunction.AddRequest) (*mathFunction.AddResponse, error) {
	log.Printf("%s\n", context)
	log.Printf("%s\n", req)
	return &mathFunction.AddResponse{
		Result: req.P1 + req.P2 + req.P3,
	}, nil
}
func (MathFunction) Subtraction(context context.Context, req *mathFunction.SubtractionRequest) (*mathFunction.SubtractionResponse, error) {
	log.Printf("%s\n", context)
	log.Printf("%s\n", req)
	return &mathFunction.SubtractionResponse{
		Result: req.P1 - req.P2,
	}, nil
}
