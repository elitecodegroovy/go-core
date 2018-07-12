package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/elitecodegroovy/go-core/chart9/go-micro/simple/proto"
	"log"
	"time"
)

//添加环境变量：
//CONSUL_HTTP_ADDR=192.168.1.229:8500
func main() {
	// 创建一个服务实例，可以指定有关的设置项
	service := micro.NewService(micro.Name("grpc.client"))
	service.Init()

	//创建一个调用者服务，grpc是注册服务的名称
	grpcCaller := proto.NewGRPCService("grpc", service.Client())

	// 调用rpc函数
	t1 := time.Now();
	rsp, err := grpcCaller.CallGRPC(context.TODO(), &proto.GRPCReq{Name: "John", ReqSeq: "000120180712"})
	if err != nil {
		log.Fatal(err)
	}else {
		//输出成功的请求
		log.Output(2, fmt.Sprintf("rpc结果: %s, 消息: %s, 耗时：%dms", rsp.Result, rsp.Msg,
			time.Since(t1).Nanoseconds()/1000000))
	}
}