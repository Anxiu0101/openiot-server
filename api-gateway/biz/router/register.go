// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	device "github.com/OpenIoT-Hub/openiot-server/api-gateway/biz/router/device"
	user "github.com/OpenIoT-Hub/openiot-server/api-gateway/biz/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	device.Register(r)

	user.Register(r)

}
