// Code generated by Kitex v0.5.2. DO NOT EDIT.

package openiotuserservice

import (
	"context"
	"fmt"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/api"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/api/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return openiotUserServiceServiceInfo
}

var openiotUserServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OpeniotUserService"
	handlerType := (*user.OpeniotUserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Ping":         kitex.NewMethodInfo(pingHandler, newPingArgs, newPingResult, false),
		"CreateUser":   kitex.NewMethodInfo(createUserHandler, newCreateUserArgs, newCreateUserResult, false),
		"RemoveUser":   kitex.NewMethodInfo(removeUserHandler, newRemoveUserArgs, newRemoveUserResult, false),
		"UpdateUser":   kitex.NewMethodInfo(updateUserHandler, newUpdateUserArgs, newUpdateUserResult, false),
		"GetUserInfo":  kitex.NewMethodInfo(getUserInfoHandler, newGetUserInfoArgs, newGetUserInfoResult, false),
		"ListUserInfo": kitex.NewMethodInfo(listUserInfoHandler, newListUserInfoArgs, newListUserInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "idl",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.PingReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.OpeniotUserService).Ping(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PingArgs:
		success, err := handler.(user.OpeniotUserService).Ping(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PingResult)
		realResult.Success = success
	}
	return nil
}
func newPingArgs() interface{} {
	return &PingArgs{}
}

func newPingResult() interface{} {
	return &PingResult{}
}

type PingArgs struct {
	Req *api.PingReq
}

func (p *PingArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.PingReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PingArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PingArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PingArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PingArgs) Unmarshal(in []byte) error {
	msg := new(api.PingReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PingArgs_Req_DEFAULT *api.PingReq

func (p *PingArgs) GetReq() *api.PingReq {
	if !p.IsSetReq() {
		return PingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PingArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PingArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PingResult struct {
	Success *api.BaseRsp
}

var PingResult_Success_DEFAULT *api.BaseRsp

func (p *PingResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.BaseRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PingResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PingResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PingResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PingResult) Unmarshal(in []byte) error {
	msg := new(api.BaseRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PingResult) GetSuccess() *api.BaseRsp {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PingResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.BaseRsp)
}

func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PingResult) GetResult() interface{} {
	return p.Success
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.CreateUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.OpeniotUserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateUserArgs:
		success, err := handler.(user.OpeniotUserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *user.CreateUserReq
}

func (p *CreateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.CreateUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(user.CreateUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *user.CreateUserReq

func (p *CreateUserArgs) GetReq() *user.CreateUserReq {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateUserResult struct {
	Success *user.CreateUserRsp
}

var CreateUserResult_Success_DEFAULT *user.CreateUserRsp

func (p *CreateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.CreateUserRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(user.CreateUserRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *user.CreateUserRsp {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.CreateUserRsp)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateUserResult) GetResult() interface{} {
	return p.Success
}

func removeUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.RemoveUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.OpeniotUserService).RemoveUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RemoveUserArgs:
		success, err := handler.(user.OpeniotUserService).RemoveUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RemoveUserResult)
		realResult.Success = success
	}
	return nil
}
func newRemoveUserArgs() interface{} {
	return &RemoveUserArgs{}
}

func newRemoveUserResult() interface{} {
	return &RemoveUserResult{}
}

type RemoveUserArgs struct {
	Req *user.RemoveUserReq
}

func (p *RemoveUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.RemoveUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RemoveUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RemoveUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RemoveUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RemoveUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RemoveUserArgs) Unmarshal(in []byte) error {
	msg := new(user.RemoveUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RemoveUserArgs_Req_DEFAULT *user.RemoveUserReq

func (p *RemoveUserArgs) GetReq() *user.RemoveUserReq {
	if !p.IsSetReq() {
		return RemoveUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RemoveUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RemoveUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RemoveUserResult struct {
	Success *user.RemoveUserRsp
}

var RemoveUserResult_Success_DEFAULT *user.RemoveUserRsp

func (p *RemoveUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.RemoveUserRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RemoveUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RemoveUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RemoveUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RemoveUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RemoveUserResult) Unmarshal(in []byte) error {
	msg := new(user.RemoveUserRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RemoveUserResult) GetSuccess() *user.RemoveUserRsp {
	if !p.IsSetSuccess() {
		return RemoveUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RemoveUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.RemoveUserRsp)
}

func (p *RemoveUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RemoveUserResult) GetResult() interface{} {
	return p.Success
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.OpeniotUserService).UpdateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateUserArgs:
		success, err := handler.(user.OpeniotUserService).UpdateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateUserArgs() interface{} {
	return &UpdateUserArgs{}
}

func newUpdateUserResult() interface{} {
	return &UpdateUserResult{}
}

type UpdateUserArgs struct {
	Req *user.UpdateUserReq
}

func (p *UpdateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserArgs_Req_DEFAULT *user.UpdateUserReq

func (p *UpdateUserArgs) GetReq() *user.UpdateUserReq {
	if !p.IsSetReq() {
		return UpdateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserResult struct {
	Success *user.UpdateUserRsp
}

var UpdateUserResult_Success_DEFAULT *user.UpdateUserRsp

func (p *UpdateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserResult) GetSuccess() *user.UpdateUserRsp {
	if !p.IsSetSuccess() {
		return UpdateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserRsp)
}

func (p *UpdateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.OpeniotUserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserInfoArgs:
		success, err := handler.(user.OpeniotUserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *user.GetUserInfoReq
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *user.GetUserInfoReq

func (p *GetUserInfoArgs) GetReq() *user.GetUserInfoReq {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *user.GetUserInfoRsp
}

var GetUserInfoResult_Success_DEFAULT *user.GetUserInfoRsp

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserInfoRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *user.GetUserInfoRsp {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserInfoRsp)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
	return p.Success
}

func listUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.ListUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.OpeniotUserService).ListUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListUserInfoArgs:
		success, err := handler.(user.OpeniotUserService).ListUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newListUserInfoArgs() interface{} {
	return &ListUserInfoArgs{}
}

func newListUserInfoResult() interface{} {
	return &ListUserInfoResult{}
}

type ListUserInfoArgs struct {
	Req *user.ListUserInfoReq
}

func (p *ListUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.ListUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.ListUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListUserInfoArgs_Req_DEFAULT *user.ListUserInfoReq

func (p *ListUserInfoArgs) GetReq() *user.ListUserInfoReq {
	if !p.IsSetReq() {
		return ListUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListUserInfoResult struct {
	Success *user.ListUserInfoRsp
}

var ListUserInfoResult_Success_DEFAULT *user.ListUserInfoRsp

func (p *ListUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.ListUserInfoRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.ListUserInfoRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListUserInfoResult) GetSuccess() *user.ListUserInfoRsp {
	if !p.IsSetSuccess() {
		return ListUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.ListUserInfoRsp)
}

func (p *ListUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListUserInfoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Ping(ctx context.Context, Req *api.PingReq) (r *api.BaseRsp, err error) {
	var _args PingArgs
	_args.Req = Req
	var _result PingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, Req *user.CreateUserReq) (r *user.CreateUserRsp, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveUser(ctx context.Context, Req *user.RemoveUserReq) (r *user.RemoveUserRsp, err error) {
	var _args RemoveUserArgs
	_args.Req = Req
	var _result RemoveUserResult
	if err = p.c.Call(ctx, "RemoveUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, Req *user.UpdateUserReq) (r *user.UpdateUserRsp, err error) {
	var _args UpdateUserArgs
	_args.Req = Req
	var _result UpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq) (r *user.GetUserInfoRsp, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListUserInfo(ctx context.Context, Req *user.ListUserInfoReq) (r *user.ListUserInfoRsp, err error) {
	var _args ListUserInfoArgs
	_args.Req = Req
	var _result ListUserInfoResult
	if err = p.c.Call(ctx, "ListUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
