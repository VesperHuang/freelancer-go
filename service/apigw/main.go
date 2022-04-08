package main

import (	
	"time"
	"github.com/micro/go-micro"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"

	"freelancer-go/service/apigw/route"
)

func main() {
    reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
		
	service := micro.NewService(
		micro.Name("go.micro.service.apigw"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Registry(reg),
	)

	service.Init()

	r := route.Router()
	r.Run(":8080")
}
