package pack

import (
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/api"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
)

func BuildBaseRsp(err int) *api.BaseRsp {
	return &api.BaseRsp{
		StatusCode: int64(err),
		StatusMsg:  errno.GetMsg(err),
	}
}
