// Code generated by Kitex v0.5.2. DO NOT EDIT.

package openiotdeviceservice

import (
	"context"
	api "github.com/OpenIoT-Hub/openiot-server/internal/device/kitex_gen/openiot/api"
	device "github.com/OpenIoT-Hub/openiot-server/internal/device/kitex_gen/openiot/api/device"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Ping(ctx context.Context, Req *device.PingReq, callOptions ...callopt.Option) (r *api.BaseRsp, err error)
	CreateDevice(ctx context.Context, Req *device.CreateDeviceReq, callOptions ...callopt.Option) (r *device.CreateDeviceRsp, err error)
	RemoveDevice(ctx context.Context, Req *device.RemoveDeviceReq, callOptions ...callopt.Option) (r *device.RemoveDeviceRsp, err error)
	UpdateDevice(ctx context.Context, Req *device.UpdateDeviceReq, callOptions ...callopt.Option) (r *device.UpdateDeviceRsp, err error)
	GetDevice(ctx context.Context, Req *device.GetDeviceReq, callOptions ...callopt.Option) (r *device.GetDeviceRsp, err error)
	ListDevice(ctx context.Context, Req *device.ListDeviceReq, callOptions ...callopt.Option) (r *device.ListDeviceRsp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kOpeniotDeviceServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kOpeniotDeviceServiceClient struct {
	*kClient
}

func (p *kOpeniotDeviceServiceClient) Ping(ctx context.Context, Req *device.PingReq, callOptions ...callopt.Option) (r *api.BaseRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Ping(ctx, Req)
}

func (p *kOpeniotDeviceServiceClient) CreateDevice(ctx context.Context, Req *device.CreateDeviceReq, callOptions ...callopt.Option) (r *device.CreateDeviceRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateDevice(ctx, Req)
}

func (p *kOpeniotDeviceServiceClient) RemoveDevice(ctx context.Context, Req *device.RemoveDeviceReq, callOptions ...callopt.Option) (r *device.RemoveDeviceRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveDevice(ctx, Req)
}

func (p *kOpeniotDeviceServiceClient) UpdateDevice(ctx context.Context, Req *device.UpdateDeviceReq, callOptions ...callopt.Option) (r *device.UpdateDeviceRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateDevice(ctx, Req)
}

func (p *kOpeniotDeviceServiceClient) GetDevice(ctx context.Context, Req *device.GetDeviceReq, callOptions ...callopt.Option) (r *device.GetDeviceRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetDevice(ctx, Req)
}

func (p *kOpeniotDeviceServiceClient) ListDevice(ctx context.Context, Req *device.ListDeviceReq, callOptions ...callopt.Option) (r *device.ListDeviceRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListDevice(ctx, Req)
}
