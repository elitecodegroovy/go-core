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

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerServiceClient, customerReq *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customerReq)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client pb.CustomerServiceClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
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
		Name:  "Liu Jigang",
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
	createCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "ZhangXiaolong",
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
	createCustomer(client, customer)
	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)
}
