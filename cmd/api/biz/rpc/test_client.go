package rpc

import (
	"context"
	"douyin/kitex_gen/test"
	"douyin/kitex_gen/test/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func Echo() {
	c, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &test.EchoReq{Message: "my request"}
	resp, err := c.Test(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
