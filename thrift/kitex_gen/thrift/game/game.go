// Code generated by Kitex v0.9.1. DO NOT EDIT.

package game

import (
	thrift0 "TestProject/thrift/kitex_gen/thrift"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetInfo": kitex.NewMethodInfo(
		getInfoHandler,
		newGameGetInfoArgs,
		newGameGetInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	gameServiceInfo                = NewServiceInfo()
	gameServiceInfoForClient       = NewServiceInfoForClient()
	gameServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return gameServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return gameServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return gameServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Game"
	handlerType := (*thrift0.Game)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "thrift",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func getInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*thrift0.GameGetInfoArgs)
	realResult := result.(*thrift0.GameGetInfoResult)
	success, err := handler.(thrift0.Game).GetInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGameGetInfoArgs() interface{} {
	return thrift0.NewGameGetInfoArgs()
}

func newGameGetInfoResult() interface{} {
	return thrift0.NewGameGetInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetInfo(ctx context.Context, req *thrift0.GReq) (r *thrift0.GResp, err error) {
	var _args thrift0.GameGetInfoArgs
	_args.Req = req
	var _result thrift0.GameGetInfoResult
	if err = p.c.Call(ctx, "GetInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
