package main

import (
	"TestProject/api/client"
	"TestProject/api/kitex_gen/pb"
	"context"
	"fmt"
)

func main() {
	cli := client.NewTestServiceClient()
	resp, _ := cli.GetInfo(context.Background(), &pb.InfoReq{Id: 1})
	fmt.Println(resp)
	resp2, _ := cli.GetInfo(context.Background(), &pb.InfoReq{Id: 2})
	fmt.Println(resp2)
}
