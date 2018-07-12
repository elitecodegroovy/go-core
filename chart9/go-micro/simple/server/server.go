package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/elitecodegroovy/go-core/chart9/go-micro/simple/proto"
	"log"
)

type GRPC struct{}

func (g *GRPC) CallGRPC(ctx context.Context, req *proto.GRPCReq, rsp *proto.GRPCResp) error {
	log.Output(2, fmt.Sprintf("req sequece NO.:%s", req.ReqSeq))
	rsp.Msg = "请求的名称： " + req.Name
	rsp.Result = "成功！"
	return nil
}

//添加环境变量：
//CONSUL_HTTP_ADDR=192.168.1.229:8500, 此地址为registry的默认consul注册地址
func main() {
	// 创建一个服务实例，可以指定有关的设置项
	service := micro.NewService(
		micro.Name("grpc"),
	)

	// 初始化解析命令行标记
	service.Init()

	// 注册服务操作
	proto.RegisterGRPCHandler(service.Server(), new(GRPC))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}