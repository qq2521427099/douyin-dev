// Code generated by Kitex v0.3.2. DO NOT EDIT.

package echo

import (
	"context"
	"douyin/kitex_gen/test"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return echoServiceInfo
}

var echoServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Echo"
	handlerType := (*test.Echo)(nil)
	methods := map[string]kitex.MethodInfo{
		"Test": kitex.NewMethodInfo(testHandler, newEchoTestArgs, newEchoTestResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "test",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func testHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*test.EchoTestArgs)
	realResult := result.(*test.EchoTestResult)
	success, err := handler.(test.Echo).Test(ctx, realArg.EchoReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEchoTestArgs() interface{} {
	return test.NewEchoTestArgs()
}

func newEchoTestResult() interface{} {
	return test.NewEchoTestResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Test(ctx context.Context, echoReq *test.EchoReq) (r *test.EchoResp, err error) {
	var _args test.EchoTestArgs
	_args.EchoReq = echoReq
	var _result test.EchoTestResult
	if err = p.c.Call(ctx, "Test", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}