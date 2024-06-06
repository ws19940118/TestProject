package client

import (
	"TestProject/api/kitex_gen/pb/testservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"time"
)

// new一个客户端，业务上可以直接调用方法，testservice.Client.GetInfo()
func NewTestServiceClient() testservice.Client {
	/*r, err := etcd.NewEtcdResolver([]string{"192.168.1.12:100,192.168.1.13:101"})
	if err != nil {
		fmt.Println(err)
	}
	fp := retry.NewFailurePolicy()
	fp.WithMaxRetryTimes(3)                          // 配置最多重试3次
	cli, err := testservice.NewClient("TestService", // 服务名
		client.WithRPCTimeout(time.Second*3), // 设置超时时间
		client.WithFailureRetry(fp),          // 设置失败重试
		client.WithResolver(r),               // 设置注册中心
		client.WithMuxConnection(1))          // 设置连接数
	if err != nil {
		fmt.Println(err)
	}
	return cli*/
	fp := retry.NewFailurePolicy()
	fp.WithMaxRetryTimes(3)                          // 配置最多重试3次
	cli, err := testservice.NewClient("TestService", // 服务名
		client.WithHostPorts("172.16.14.244:8081"), // 服务地址
		client.WithRPCTimeout(time.Second*3),       // 设置超时时间
		client.WithFailureRetry(fp),                // 设置失败重试
		client.WithMuxConnection(1))                // 设置连接数
	if err != nil {
		fmt.Println(err)
	}
	return cli
}
