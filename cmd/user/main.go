package main

import (
	"douyin/cmd/user/conn"
	user "douyin/kitex_gen/user/userservice"
	"douyin/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {

	conn.InitGorm()

	addr, _ := net.ResolveTCPAddr("tcp", conf.UserPort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := user.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
