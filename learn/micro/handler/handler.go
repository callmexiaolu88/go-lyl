package handler

import (
	"context"
	mathFunction "lyl/learn/micro/proto"
)

type MathFunction struct{}

func (*MathFunction) Add(context context.Context, req *mathFunction.AddRequest, rsp *mathFunction.AddResponse) error {
	rsp.Result = req.P1 + req.P2 + req.P3
	return nil
}

func (*MathFunction) Subtraction(context context.Context, req *mathFunction.SubtractionRequest, rsp *mathFunction.SubtractionResponse) error {
	rsp.Result = req.P1 - req.P2
	return nil
}
