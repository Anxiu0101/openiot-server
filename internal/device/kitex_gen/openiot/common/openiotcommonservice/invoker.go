// Code generated by Kitex v0.4.4. DO NOT EDIT.

package openiotcommonservice

import (
	common "github.com/OpenIoT-Hub/openiot-server/internal/device/kitex_gen/openiot/common"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler common.OpeniotCommonService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
