package main

import (
	"context"
	"douyin/kitex_gen/test"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Test implements the EchoImpl interface.
func (s *EchoImpl) Test(ctx context.Context, echoReq *test.EchoReq) (resp *test.EchoResp, err error) {
	// TODO: Your code here...
	resp = &test.EchoResp{
		BaseResp: &test.BaseResp{StatusCode: 0},
		Message:  echoReq.Message,
	}
	return
}
