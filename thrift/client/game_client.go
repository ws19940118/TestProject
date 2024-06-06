package client

import (
	"TestProject/thrift/kitex_gen/thrift/game"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"time"
)

func NewGameClient() game.Client {
	fp := retry.NewFailurePolicy()
	fp.WithMaxRetryTimes(3)                   // 配置最多重试3次
	cli, err := game.NewClient("TestService", // 服务名
		client.WithHostPorts("172.16.14.244:8081"), // 服务地址
		client.WithRPCTimeout(time.Second*3),       // 设置超时时间
		client.WithFailureRetry(fp),                // 设置失败重试
		client.WithMuxConnection(1))                // 设置连接数
	if err != nil {
		fmt.Println(err)
	}
	return cli
}
