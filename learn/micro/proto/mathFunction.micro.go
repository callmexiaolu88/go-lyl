// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mathFunction.proto

/*
Package mathFunction is a generated protocol buffer package.

It is generated from these files:
	mathFunction.proto

It has these top-level messages:
	AddRequest
	AddResponse
	SubtractionRequest
	SubtractionResponse
*/
package mathFunction

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MathFunction service

type MathFunctionService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Subtraction(ctx context.Context, in *SubtractionRequest, opts ...client.CallOption) (*SubtractionResponse, error)
}

type mathFunctionService struct {
	c    client.Client
	name string
}

func NewMathFunctionService(name string, c client.Client) MathFunctionService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mathfunction"
	}
	return &mathFunctionService{
		c:    c,
		name: name,
	}
}

func (c *mathFunctionService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "MathFunction.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathFunctionService) Subtraction(ctx context.Context, in *SubtractionRequest, opts ...client.CallOption) (*SubtractionResponse, error) {
	req := c.c.NewRequest(c.name, "MathFunction.Subtraction", in)
	out := new(SubtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MathFunction service

type MathFunctionHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Subtraction(context.Context, *SubtractionRequest, *SubtractionResponse) error
}

func RegisterMathFunctionHandler(s server.Server, hdlr MathFunctionHandler, opts ...server.HandlerOption) error {
	type mathFunction interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Subtraction(ctx context.Context, in *SubtractionRequest, out *SubtractionResponse) error
	}
	type MathFunction struct {
		mathFunction
	}
	h := &mathFunctionHandler{hdlr}
	return s.Handle(s.NewHandler(&MathFunction{h}, opts...))
}

type mathFunctionHandler struct {
	MathFunctionHandler
}

func (h *mathFunctionHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.MathFunctionHandler.Add(ctx, in, out)
}

func (h *mathFunctionHandler) Subtraction(ctx context.Context, in *SubtractionRequest, out *SubtractionResponse) error {
	return h.MathFunctionHandler.Subtraction(ctx, in, out)
}
