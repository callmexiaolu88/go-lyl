// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: water.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	water.proto

It has these top-level messages:
	Empty
	DeviceInfo
	DeviceCountConditional
	DeviceCountResponse
	DeviceCountsConditional
	DeviceCountsResponse
	DeviceValueConditional
	DeviceValueResponse
	DeviceValuesConditional
	DeviceValuesResponse
	DeviceDetail
	DeviceDetailConditional
	DeviceDetailResponse
	DeviceDetailsConditional
	DeviceDetailsResponse
*/
package main

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

// Client API for WaterService service

type WaterService interface {
	GetDeviceCount(ctx context.Context, in *DeviceCountConditional, opts ...client.CallOption) (*DeviceCountResponse, error)
	GetDeviceCounts(ctx context.Context, in *DeviceCountsConditional, opts ...client.CallOption) (*DeviceCountsResponse, error)
	GetDeviceValue(ctx context.Context, in *DeviceValueConditional, opts ...client.CallOption) (*DeviceValueResponse, error)
	GetDeviceValues(ctx context.Context, in *DeviceValuesConditional, opts ...client.CallOption) (*DeviceValuesResponse, error)
	GetDeviceDetail(ctx context.Context, in *DeviceDetailConditional, opts ...client.CallOption) (*DeviceDetailResponse, error)
	GetDeviceDetails(ctx context.Context, in *DeviceDetailsConditional, opts ...client.CallOption) (*DeviceDetailsResponse, error)
}

type waterService struct {
	c    client.Client
	name string
}

func NewWaterService(name string, c client.Client) WaterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "main"
	}
	return &waterService{
		c:    c,
		name: name,
	}
}

func (c *waterService) GetDeviceCount(ctx context.Context, in *DeviceCountConditional, opts ...client.CallOption) (*DeviceCountResponse, error) {
	req := c.c.NewRequest(c.name, "WaterService.GetDeviceCount", in)
	out := new(DeviceCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waterService) GetDeviceCounts(ctx context.Context, in *DeviceCountsConditional, opts ...client.CallOption) (*DeviceCountsResponse, error) {
	req := c.c.NewRequest(c.name, "WaterService.GetDeviceCounts", in)
	out := new(DeviceCountsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waterService) GetDeviceValue(ctx context.Context, in *DeviceValueConditional, opts ...client.CallOption) (*DeviceValueResponse, error) {
	req := c.c.NewRequest(c.name, "WaterService.GetDeviceValue", in)
	out := new(DeviceValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waterService) GetDeviceValues(ctx context.Context, in *DeviceValuesConditional, opts ...client.CallOption) (*DeviceValuesResponse, error) {
	req := c.c.NewRequest(c.name, "WaterService.GetDeviceValues", in)
	out := new(DeviceValuesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waterService) GetDeviceDetail(ctx context.Context, in *DeviceDetailConditional, opts ...client.CallOption) (*DeviceDetailResponse, error) {
	req := c.c.NewRequest(c.name, "WaterService.GetDeviceDetail", in)
	out := new(DeviceDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waterService) GetDeviceDetails(ctx context.Context, in *DeviceDetailsConditional, opts ...client.CallOption) (*DeviceDetailsResponse, error) {
	req := c.c.NewRequest(c.name, "WaterService.GetDeviceDetails", in)
	out := new(DeviceDetailsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WaterService service

type WaterServiceHandler interface {
	GetDeviceCount(context.Context, *DeviceCountConditional, *DeviceCountResponse) error
	GetDeviceCounts(context.Context, *DeviceCountsConditional, *DeviceCountsResponse) error
	GetDeviceValue(context.Context, *DeviceValueConditional, *DeviceValueResponse) error
	GetDeviceValues(context.Context, *DeviceValuesConditional, *DeviceValuesResponse) error
	GetDeviceDetail(context.Context, *DeviceDetailConditional, *DeviceDetailResponse) error
	GetDeviceDetails(context.Context, *DeviceDetailsConditional, *DeviceDetailsResponse) error
}

func RegisterWaterServiceHandler(s server.Server, hdlr WaterServiceHandler, opts ...server.HandlerOption) error {
	type waterService interface {
		GetDeviceCount(ctx context.Context, in *DeviceCountConditional, out *DeviceCountResponse) error
		GetDeviceCounts(ctx context.Context, in *DeviceCountsConditional, out *DeviceCountsResponse) error
		GetDeviceValue(ctx context.Context, in *DeviceValueConditional, out *DeviceValueResponse) error
		GetDeviceValues(ctx context.Context, in *DeviceValuesConditional, out *DeviceValuesResponse) error
		GetDeviceDetail(ctx context.Context, in *DeviceDetailConditional, out *DeviceDetailResponse) error
		GetDeviceDetails(ctx context.Context, in *DeviceDetailsConditional, out *DeviceDetailsResponse) error
	}
	type WaterService struct {
		waterService
	}
	h := &waterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WaterService{h}, opts...))
}

type waterServiceHandler struct {
	WaterServiceHandler
}

func (h *waterServiceHandler) GetDeviceCount(ctx context.Context, in *DeviceCountConditional, out *DeviceCountResponse) error {
	return h.WaterServiceHandler.GetDeviceCount(ctx, in, out)
}

func (h *waterServiceHandler) GetDeviceCounts(ctx context.Context, in *DeviceCountsConditional, out *DeviceCountsResponse) error {
	return h.WaterServiceHandler.GetDeviceCounts(ctx, in, out)
}

func (h *waterServiceHandler) GetDeviceValue(ctx context.Context, in *DeviceValueConditional, out *DeviceValueResponse) error {
	return h.WaterServiceHandler.GetDeviceValue(ctx, in, out)
}

func (h *waterServiceHandler) GetDeviceValues(ctx context.Context, in *DeviceValuesConditional, out *DeviceValuesResponse) error {
	return h.WaterServiceHandler.GetDeviceValues(ctx, in, out)
}

func (h *waterServiceHandler) GetDeviceDetail(ctx context.Context, in *DeviceDetailConditional, out *DeviceDetailResponse) error {
	return h.WaterServiceHandler.GetDeviceDetail(ctx, in, out)
}

func (h *waterServiceHandler) GetDeviceDetails(ctx context.Context, in *DeviceDetailsConditional, out *DeviceDetailsResponse) error {
	return h.WaterServiceHandler.GetDeviceDetails(ctx, in, out)
}
