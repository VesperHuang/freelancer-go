
package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
    // "github.com/micro/go-micro/registry"
    // "github.com/micro/go-plugins/registry/consul"

	"freelancer-go/common"
	"freelancer-go/service/account/handler"
	proto "freelancer-go/service/account/proto"
	dbproxy "freelancer-go/service/dbproxy/client"
)

func main() {
	// 创建 consul 服务注册项，其中 192.168.3.25:2379 为 consul 服务地址。
    // consul 服务地址按照实际情况填写
    // reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
		
	service := micro.NewService(
		// service := k8s.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		// micro.Registry(reg),
	)

	// 初始化service, 解析命令行参数等
	service.Init()
	
	// 初始化dbproxy client
	dbproxy.Init(service)

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
