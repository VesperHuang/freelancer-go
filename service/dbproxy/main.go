package main

import (
	"freelancer-go/common"
	"freelancer-go/service/dbproxy/config"
	"log"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"	

	// _ "github.com/micro/go-plugins/registry/kubernetes"

	dbConn "freelancer-go/service/dbproxy/conn"
	dbProxy "freelancer-go/service/dbproxy/proto"
	dbRpc "freelancer-go/service/dbproxy/rpc"
)

func startRpcService() {
	
	// 创建 consul 服务注册项，其中 192.168.3.25:2379 为 consul 服务地址。
    // consul 服务地址按照实际情况填写
    reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	
	// reg := consul.NewRegistry(func(options *registry.Options) {
	// 	options.Addrs = []string{
	// 		"127.0.0.1:8300",
	// 	}
	// })

	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"), // 在注册中心中的服务名称
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*10),      // 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(reg),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 检查是否指定dbhost
			dbhost := c.String("dbhost")
			if len(dbhost) > 0 {
				log.Println("custom db address: " + dbhost)
				config.UpdateDBHost(dbhost)
			}
		}),
	)

	// 初始化db connection
	dbConn.InitDBConn()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
}

// res, err := mapper.FuncCall("/user/UserExist", []interface{}{"haha"}...)
// log.Printf("error: %+v\n", err)
// log.Printf("result: %+v\n", res[0].Interface())

// res, err = mapper.FuncCall("/user/UserExist", []interface{}{"admin"}...)
// log.Printf("error: %+v\n", err)
// log.Printf("result: %+v\n", res[0].Interface())
