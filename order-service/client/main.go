package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fadilrifqi/order-service/proto" // Import the generated proto package

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call CreateOrder
	createOrderResp, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
		OrderId:    "123",
		Price: 	100.0,
	})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}
	log.Printf("CreateOrder Response: OrderID=%s, Status=%s", createOrderResp.OrderId, createOrderResp.Status)

	// Call GetOrderStatus
	getOrderStatusResp, err := client.GetOrderStatus(ctx, &pb.GetOrderStatusRequest{
		OrderId: "123",
	})
	if err != nil {
		log.Fatalf("could not get order status: %v", err)
	}
	log.Printf("GetOrderStatus Response: OrderID=%s, Status=%s", getOrderStatusResp.OrderId, getOrderStatusResp.Status)

}
