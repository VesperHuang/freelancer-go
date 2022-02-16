package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"

	"freelancer-go/common"
	"freelancer-go/service/account/handler"
	proto "freelancer-go/service/account/proto"
	dbproxy "freelancer-go/service/dbproxy/client"
)

func main() {
	service := micro.NewService(
		// service := k8s.NewService(
		micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
	)

	// 初始化service, 解析命令行参数等
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	proto.RegisterUserServiceHandler(service.Server(), new(handler.UserServiceHandler))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
