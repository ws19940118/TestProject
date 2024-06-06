package main

import (
	"TestProject/thrift/client"
	"TestProject/thrift/kitex_gen/thrift"
	"context"
	"fmt"
)

func main() {
	cli := client.NewGameClient()
	resp, _ := cli.GetInfo(context.Background(), &thrift.GReq{User: &thrift.User{
		Name: "1",
		Age:  9,
	}})
	fmt.Println(resp)
	resp2, _ := cli.GetInfo(context.Background(), &thrift.GReq{User: &thrift.User{
		Name: "1",
		Age:  90,
	}})
	fmt.Println(resp2)

}
