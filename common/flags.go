package common

import "github.com/micro/cli"

// CustomFlags : 自定义命令行参数
var CustomFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "dbhost",
		Value: "172.22.19.39:13306",
		Usage: "database address",
	},
	cli.StringFlag{
		Name:  "mqhost",
		Value: "127.0.0.1",
		Usage: "mq(rabbitmq) address",
	},
	cli.StringFlag{
		Name:  "cachehost",
		Value: "127.0.0.1",
		Usage: "cache(redis) address",
	},
	cli.StringFlag{
		Name:  "cephhost",
		Value: "127.0.0.1",
		Usage: "ceph address",
	},
}
