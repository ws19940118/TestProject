package server

import (
	"TestProject/thrift/kitex_gen/thrift"
	"TestProject/thrift/kitex_gen/thrift/game"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"net"
	"strings"
)

func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func NewRpcServer() server.Server {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	/*	//调用testservice->server.go的实例化服务器的方法
		//server.WithServiceAdd():设置服务ip
		//server.WithMuxTransport() 启用多路复用传输
		//server.WithServerBasicInfo():设置服务名
		//server.WithRegistry(r) :启用etcd

		r, err := etcd.NewEtcdRegistry([]string{"172.16.1.201:2379"}) // r should not be reused.
		if err != nil {
			fmt.Errorf("NewEtcdRegistry : %v", err)
		}
		//addr := utils.NewNetAddr("tcp", fmt.Sprintf("%s:8081", InternalIP()))
		result := game.NewServer(new(GameImpl),
			server.WithRegistry(r),
			server.WithMuxTransport(),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "GameThrift"}))
		result.Run()
		//fmt.Println(addr.String())
		return result*/
	//调用testservice->server.go的实例化服务器的方法
	//server.WithServiceAdd():设置服务ip
	//server.WithMuxTransport() 启用多路复用传输
	//server.WithServerBasicInfo():设置服务名
	addr := utils.NewNetAddr("tcp", fmt.Sprintf("%s:8081", InternalIP()))
	result := game.NewServer(new(GameImpl),
		server.WithServiceAddr(addr),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "TestService"}))
	result.Run()
	fmt.Println(addr.String())
	return result
}

type GameImpl struct {
}

func (g *GameImpl) GetInfo(ctx context.Context, req *thrift.GReq) (r *thrift.GResp, err error) {
	if req.User.Age <= 10 {
		return &thrift.GResp{Msg: &thrift.Teacher{
			Name: "101010",
			Age:  10,
		}}, nil
	} else {
		return &thrift.GResp{Msg: &thrift.Teacher{
			Name: "111111",
			Age:  111,
		}}, nil
	}
}
