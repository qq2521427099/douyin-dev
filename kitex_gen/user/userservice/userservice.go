// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyin/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ContainsName":  kitex.NewMethodInfo(containsNameHandler, newUserServiceContainsNameArgs, newUserServiceContainsNameResult, false),
		"CreateUser":    kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"GetUserByName": kitex.NewMethodInfo(getUserByNameHandler, newUserServiceGetUserByNameArgs, newUserServiceGetUserByNameResult, false),
		"GetUserById":   kitex.NewMethodInfo(getUserByIdHandler, newUserServiceGetUserByIdArgs, newUserServiceGetUserByIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func containsNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceContainsNameArgs)
	realResult := result.(*user.UserServiceContainsNameResult)
	success, err := handler.(user.UserService).ContainsName(ctx, realArg.ContainsNameReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceContainsNameArgs() interface{} {
	return user.NewUserServiceContainsNameArgs()
}

func newUserServiceContainsNameResult() interface{} {
	return user.NewUserServiceContainsNameResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCreateUserArgs)
	realResult := result.(*user.UserServiceCreateUserResult)
	success, err := handler.(user.UserService).CreateUser(ctx, realArg.CreateUserReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return user.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return user.NewUserServiceCreateUserResult()
}

func getUserByNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserByNameArgs)
	realResult := result.(*user.UserServiceGetUserByNameResult)
	success, err := handler.(user.UserService).GetUserByName(ctx, realArg.GetUserByNameReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserByNameArgs() interface{} {
	return user.NewUserServiceGetUserByNameArgs()
}

func newUserServiceGetUserByNameResult() interface{} {
	return user.NewUserServiceGetUserByNameResult()
}

func getUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserByIdArgs)
	realResult := result.(*user.UserServiceGetUserByIdResult)
	success, err := handler.(user.UserService).GetUserById(ctx, realArg.GetUserByIdReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserByIdArgs() interface{} {
	return user.NewUserServiceGetUserByIdArgs()
}

func newUserServiceGetUserByIdResult() interface{} {
	return user.NewUserServiceGetUserByIdResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ContainsName(ctx context.Context, containsNameReq *user.ContainsNameReq) (r *user.ContainsNameResp, err error) {
	var _args user.UserServiceContainsNameArgs
	_args.ContainsNameReq = containsNameReq
	var _result user.UserServiceContainsNameResult
	if err = p.c.Call(ctx, "ContainsName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, createUserReq *user.CreateUserReq) (r *user.CreateUserResp, err error) {
	var _args user.UserServiceCreateUserArgs
	_args.CreateUserReq = createUserReq
	var _result user.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserByName(ctx context.Context, getUserByNameReq *user.GetUserByNameReq) (r *user.GetUserByNameResp, err error) {
	var _args user.UserServiceGetUserByNameArgs
	_args.GetUserByNameReq = getUserByNameReq
	var _result user.UserServiceGetUserByNameResult
	if err = p.c.Call(ctx, "GetUserByName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserById(ctx context.Context, getUserByIdReq *user.GetUserByIdReq) (r *user.GetUserByIdResp, err error) {
	var _args user.UserServiceGetUserByIdArgs
	_args.GetUserByIdReq = getUserByIdReq
	var _result user.UserServiceGetUserByIdResult
	if err = p.c.Call(ctx, "GetUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
