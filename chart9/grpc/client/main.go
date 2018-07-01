package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/elitecodegroovy/go-core/chart9/grpc/customer"
)

const (
	address = "localhost:50051"
)

// 通过调用RPC方法CreateCustomer添加消费者
func addCustomer(client pb.CustomerServiceClient, customerReq *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customerReq)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// 通过指定的过滤条件CustomerFilter作为参数来查询消费者信息
func queryCustomers(client pb.CustomerServiceClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("RPC GetCustomers 异常: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("消费者：: %v", customer)
		log.Printf("首选地址的所在城市: %v", customer.Addresses[0].City)
	}
}

//程序入口
func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerServiceClient(conn)

	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "刘继刚",
		Email: "elite_jigang@163.com",
		Phone: "15914313549",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "天河区体育中心中信大厦",
				City:              "广州市",
				State:             "广东省",
				Zip:               "12516",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "天河区珠江新城IFC",
				City:              "广州市",
				State:             "广东省",
				Zip:               "12517",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	addCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "李强",
		Email: "zhangXiaolong@tencent.com",
		Phone: "139155988",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "天河区珠江新城",
				City:               "广州市",
				State:              "广东省",
				Zip:               "12519",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	addCustomer(client, customer)
	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: ""}
	queryCustomers(client, filter)
}
