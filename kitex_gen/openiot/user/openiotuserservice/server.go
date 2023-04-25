// Code generated by Kitex v0.4.4. DO NOT EDIT.
package openiotuserservice

import (
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user.OpeniotUserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
