package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/elitecodegroovy/go-core/chart9/grpc/customer"
	"fmt"
)

const (
	port = ":50051"
)

// server is used to implement customer.CustomerServiceServer.
type server struct {
	savedCustomers []*pb.CustomerRequest
}


func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	log.Output(2, fmt.Sprintf("req: %v", in))
	//类型为CustomerRequest的数组保存请求的Customer对象
	s.savedCustomers = append(s.savedCustomers, in)
	//TODO ...可以将数据持久化到关系型数据库或者NoSQL数据库.
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// 对接口CustomerServiceServer方法GetCustomers的实现。
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.CustomerService_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		//以流的形式发送给客户端，客户端也使用流的模式进行接受。
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	//指定我们想监听的端口
	lis, err := net.Listen("tcp", port)
	//端口被占用或者其它可能的异常原因
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个RPC服务器
	s := grpc.NewServer()
	//注册服务
	pb.RegisterCustomerServiceServer(s, &server{})
	//阻塞等待直到进程被杀死或者stop()被调用
	s.Serve(lis)
}
