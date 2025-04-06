package main

import (
	"context"
	"log"
	"net"
	"sync"

	orderpb "github.com/fadilrifqi/order-service/proto"      // Import the generated proto package for OrderService
	paymentpb "github.com/fadilrifqi/payment-service/proto"  // Import the generated proto package for PaymentService
	servicepb "github.com/fadilrifqi/shipping-service/proto" // Import the generated proto package

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ShippingServiceServer is the server implementation for the ShippingService
type ShippingServiceServer struct {
    servicepb.UnimplementedShippingServiceServer
    mu             sync.Mutex
    shippingStatus map[string]string // A simple in-memory store for shipping statuses
}

// NewShippingServiceServer creates a new ShippingServiceServer
func NewShippingServiceServer() *ShippingServiceServer {
    return &ShippingServiceServer{
        shippingStatus: make(map[string]string),
    }
}

// StartShipping handles the initiation of shipping
func (s *ShippingServiceServer) StartShipping(ctx context.Context, req *servicepb.StartShippingRequest) (*servicepb.StartShippingResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Save shipping status as PENDING
    s.shippingStatus[req.OrderId] = "PENDING"
    log.Printf("Shipping started: OrderID=%s, Address=%s, Status=PENDING", req.OrderId, req.Address)

    return &servicepb.StartShippingResponse{
        OrderId: req.OrderId,
        Address: req.Address,
        Status:  "PENDING",
    }, nil
}

// CancelShipping handles the cancellation of shipping
func (s *ShippingServiceServer) CancelShipping(ctx context.Context, req *servicepb.CancelShippingRequest) (*servicepb.CancelShippingResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the shipping exists
    status, exists := s.shippingStatus[req.OrderId]
    if !exists || status != "PENDING" {
        log.Printf("Cancel failed: OrderID=%s not found or not in PENDING state", req.OrderId)
        return &servicepb.CancelShippingResponse{
            Success: false,
            Message: "Shipping not found or not in PENDING state",
        }, nil
    }

    // Update shipping status to CANCELLED
    s.shippingStatus[req.OrderId] = "CANCELLED"
    log.Printf("Shipping cancelled: OrderID=%s, Status=CANCELLED", req.OrderId)

    // Call RefundPayment in PaymentService
    if err := callRefundPayment(ctx, req.OrderId); err != nil {
        log.Printf("Failed to refund payment for OrderID=%s: %v", req.OrderId, err)
    }

    // Call CancelOrder in OrderService
    if err := callCancelOrder(ctx, req.OrderId); err != nil {
        log.Printf("Failed to cancel order for OrderID=%s: %v", req.OrderId, err)
    }

    return &servicepb.CancelShippingResponse{
        Success: true,
        Message: "Shipping cancelled successfully",
    }, nil
}

// Shipped marks the shipping as completed
func (s *ShippingServiceServer) Shipped(ctx context.Context, req *servicepb.ShippedRequest) (*servicepb.ShippedResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the shipping exists
    status, exists := s.shippingStatus[req.OrderId]
    if !exists || status != "PENDING" {
        log.Printf("Shipped failed: OrderID=%s not found or not in PENDING state", req.OrderId)
        return &servicepb.ShippedResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }

    // Update shipping status to SHIPPED
    s.shippingStatus[req.OrderId] = "SHIPPED"
    log.Printf("Shipping completed: OrderID=%s, Status=SHIPPED", req.OrderId)

    // Call CompleteOrder in OrderService
    if err := callCompleteOrder(ctx, req.OrderId); err != nil {
        log.Printf("Failed to complete order for OrderID=%s: %v", req.OrderId, err)
        return &servicepb.ShippedResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }

    return &servicepb.ShippedResponse{
        OrderId: req.OrderId,
        Status:  "SHIPPED",
    }, nil
}

// callRefundPayment calls the RefundPayment method in PaymentService
func callRefundPayment(ctx context.Context, orderId string) error {
    conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return err
    }
    defer conn.Close()

    client := paymentpb.NewPaymentServiceClient(conn)
    _, err = client.RefundPayment(ctx, &paymentpb.RefundPaymentRequest{
        OrderId: orderId,
    })
    return err
}

func callCompleteOrder(ctx context.Context, orderId string) error {
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return err
    }
    defer conn.Close()

    client := orderpb.NewOrderServiceClient(conn)
    _, err = client.CompleteOrder(ctx, &orderpb.CompleteOrderRequest{
        OrderId: orderId,
    })
    return err
}

// callCancelOrder calls the CancelOrder method in OrderService
func callCancelOrder(ctx context.Context, orderId string) error {
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return err
    }
    defer conn.Close()

    client := orderpb.NewOrderServiceClient(conn)
    _, err = client.CancelOrder(ctx, &orderpb.CancelOrderRequest{
        OrderId: orderId,
    })
    return err
}

func main() {
    // Create a listener on TCP port 50053
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the ShippingServiceServer
    shippingService := NewShippingServiceServer()
    servicepb.RegisterShippingServiceServer(grpcServer, shippingService)

    log.Println("ShippingService server is running on port 50053...")
    // Start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
