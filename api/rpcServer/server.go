package rpcServer

import (
	"TestProject/api/kitex_gen/pb"
	"TestProject/api/kitex_gen/pb/testservice"
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

// 实例化 rpc服务器对象
// 当有客户端连接过来，会触发相应的方法 TestServiceImpl.GetInfo()
// 返回--> kitex的服务器对象 server.Server
func NewTestService() server.Server {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	//调用testservice->server.go的实例化服务器的方法
	//server.WithServiceAdd():设置服务ip
	//server.WithMuxTransport() 启用多路复用传输
	//server.WithServerBasicInfo():设置服务名
	addr := utils.NewNetAddr("tcp", fmt.Sprintf("%s:8081", InternalIP()))
	result := testservice.NewServer(new(TestServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "TestService"}))
	result.Run()
	fmt.Println(addr.String())
	return result
}

// api/kitex_gen/pb/testservice/server.go
// ////重写TestService接口，服务器对象
type TestServiceImpl struct {
}

func (t *TestServiceImpl) GetInfo(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	if req.Id == 1 {
		resp = &pb.InfoResp{
			User: &pb.User{
				Name: "user1",
				Age:  "18",
				Sex:  "male",
			},
			Teacher: &pb.Teacher{
				Name: "teacher1",
				Age:  "30",
				Sex:  "female",
			},
		}
	} else {
		resp = &pb.InfoResp{
			User: &pb.User{
				Name: "user2",
				Age:  "19",
				Sex:  "female",
			},
			Teacher: &pb.Teacher{
				Name: "teacher2",
				Age:  "31",
				Sex:  "male",
			},
		}
	}
	return resp, nil
}
