package main

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	paymentpb "github.com/fadilrifqi/payment-service/proto"

	pb "github.com/fadilrifqi/order-service/proto"
)

// OrderServiceServer is the server implementation for the OrderService
type OrderServiceServer struct {
    pb.UnimplementedOrderServiceServer
    mu     sync.Mutex
    orders map[string]string // A simple in-memory store for orders and their statuses
	paymentClient paymentpb.PaymentServiceClient // gRPC client for PaymentService

}

// NewOrderServiceServer creates a new OrderServiceServer
func NewOrderServiceServer(paymentClient paymentpb.PaymentServiceClient) *OrderServiceServer {
    return &OrderServiceServer{
        orders:       make(map[string]string),
        paymentClient: paymentClient,
    }
}

// CreateOrder handles the creation of an order and initializes its status
func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Save the order with a "PENDING" status
    s.orders[req.OrderId] = "PENDING"
    log.Printf("Order created: ID=%s, Status=PENDING", req.OrderId)

    // Call CreatePayment in PaymentService
    paymentResp, err := s.paymentClient.CreatePayment(ctx, &paymentpb.CreatePaymentRequest{
        OrderId: req.OrderId,
		Price:   req.Price,
    })
    if err != nil {
        log.Printf("Failed to create payment for OrderID=%s: %v", req.OrderId, err)

        // Update the order status to "CANCELLED" or "FAILED"
        s.orders[req.OrderId] = "FAILED"
        return &pb.CreateOrderResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }

    log.Printf("Payment created successfully for OrderID=%s: Status=%s", paymentResp.OrderId, paymentResp.Status)

    // Update the order status based on the payment response
    s.orders[req.OrderId] = paymentResp.Status
    return &pb.CreateOrderResponse{
        OrderId: req.OrderId,
        Status:  paymentResp.Status,
    }, nil
}

// GetOrderStatus retrieves the status of an order
func (s *OrderServiceServer) GetOrderStatus(ctx context.Context, req *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the order exists
    status, exists := s.orders[req.OrderId]
    if !exists {
        log.Printf("Order not found: ID=%s", req.OrderId)
        return &pb.GetOrderStatusResponse{
            OrderId: req.OrderId,
            Status:  "NOT_FOUND",
        }, nil
    }

    log.Printf("Order status retrieved: ID=%s, Status=%s", req.OrderId, status)
    return &pb.GetOrderStatusResponse{
        OrderId: req.OrderId,
        Status:  status,
    }, nil
}

// CancelOrder handles the cancellation of an order
func (s *OrderServiceServer) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the order exists
    status, exists := s.orders[req.OrderId]
    if !exists {
        log.Printf("Order not found for cancellation: ID=%s", req.OrderId)
        return &pb.CancelOrderResponse{
            OrderId: req.OrderId,
            Status: "NOT_FOUND",
        }, nil
    }

    // Update the order status to "CANCELLED"
    s.orders[req.OrderId] = "CANCELLED"
    log.Printf("Order cancelled: ID=%s, Previous Status=%s, New Status=CANCELLED", req.OrderId, status)

    return &pb.CancelOrderResponse{
        OrderId: req.OrderId,
        Status:  "CANCELLED",
    }, nil
}

func (s *OrderServiceServer) CompleteOrder(ctx context.Context, req *pb.CompleteOrderRequest) (*pb.CompleteOrderResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the order exists
    status, exists := s.orders[req.OrderId]
    if !exists {
        log.Printf("Order not found for cancellation: ID=%s", req.OrderId)
        return &pb.CompleteOrderResponse{
            OrderId: req.OrderId,
            Status: "NOT_FOUND",
        }, nil
    }

    // Update the order status to "CANCELLED"
    s.orders[req.OrderId] = "COMPLETED"
    log.Printf("Order completed: ID=%s, Previous Status=%s, New Status=COMPLETED", req.OrderId, status)

    return &pb.CompleteOrderResponse{
        OrderId: req.OrderId,
        Status:  "COMPLETED",
    }, nil
}







func main() {
    paymentConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to PaymentService: %v", err)
    }
    defer paymentConn.Close()

    paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)

    // Create a listener on TCP port 50051
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the OrderServiceServer
    orderService := NewOrderServiceServer(paymentClient)
    pb.RegisterOrderServiceServer(grpcServer, orderService)

    log.Println("OrderService server is running on port 50051...")
    // Start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
