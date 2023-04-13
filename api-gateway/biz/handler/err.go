package handler

import (
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type ErrResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func SendErrorResponse(c *app.RequestContext, statusCode int64) {
	c.JSON(consts.StatusOK, ErrResponse{
		Code: statusCode,
		Msg:  errno.GetMsg(int(statusCode)),
	})
}

func SendCommonResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}

// PhaseToken returns Userid
func PhaseToken(token string) (int64, error) {
	//return rpc.CheckToken(context.Background(), &user.CheckTokenRequest{
	//	Token: token,
	//})
	return 0, nil
}
