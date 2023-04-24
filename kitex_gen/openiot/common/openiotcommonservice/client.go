// Code generated by Kitex v0.4.4. DO NOT EDIT.

package openiotcommonservice

import (
	"context"
	common "github.com/OpenIoT-Hub/openiot-server/internal/user/kitex_gen/openiot/common"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Ping(ctx context.Context, Req *common.PingReq, callOptions ...callopt.Option) (r *common.BaseRsp, err error)
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
	return &kOpeniotCommonServiceClient{
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

type kOpeniotCommonServiceClient struct {
	*kClient
}

func (p *kOpeniotCommonServiceClient) Ping(ctx context.Context, Req *common.PingReq, callOptions ...callopt.Option) (r *common.BaseRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Ping(ctx, Req)
}
