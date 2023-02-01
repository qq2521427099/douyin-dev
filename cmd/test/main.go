package main

import (
	test "douyin/kitex_gen/test/echo"
	"log"
)

func main() {
	svr := test.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
