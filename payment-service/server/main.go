package main

import (
	"context"
	"log"
	"net"
	"sync"

	orderpb "github.com/fadilrifqi/order-service/proto"       // Import the generated proto package for OrderService
	pb "github.com/fadilrifqi/payment-service/proto"          // Import the generated proto package
	shippingpb "github.com/fadilrifqi/shipping-service/proto" // Import the generated proto package for ShippingService

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PaymentServiceServer is the server implementation for the PaymentService
type PaymentServiceServer struct {
    pb.UnimplementedPaymentServiceServer
    mu          sync.Mutex
    paymentData map[string]string // A simple in-memory store for payment statuses
}

// NewPaymentServiceServer creates a new PaymentServiceServer
func NewPaymentServiceServer() *PaymentServiceServer {
    return &PaymentServiceServer{
        paymentData: make(map[string]string),
    }
}

// CreatePayment handles payment creation
func (s *PaymentServiceServer) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Save payment status as PENDING
    s.paymentData[req.OrderId] = "PENDING"
    log.Printf("Payment created: OrderID=%s, Price=%.2f, Status=PENDING", req.OrderId, req.Price)

    return &pb.CreatePaymentResponse{
        OrderId: req.OrderId,
        Status:  "PENDING",
    }, nil
}

// CancelPayment handles payment cancellation
func (s *PaymentServiceServer) CancelPayment(ctx context.Context, req *pb.CancelPaymentRequest) (*pb.CancelPaymentResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the payment exists
    status, exists := s.paymentData[req.OrderId]
    if !exists || status != "PENDING" {
        log.Printf("Cancel failed: OrderID=%s not found or not in PENDING state", req.OrderId)
        return &pb.CancelPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }

    // Update payment status to CANCELLED
    s.paymentData[req.OrderId] = "CANCELLED"
    log.Printf("Payment cancelled: OrderID=%s, Status=CANCELLED", req.OrderId)

    // Call CancelOrder in OrderService
    orderConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Printf("Failed to connect to OrderService: %v", err)
        return &pb.CancelPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }
    defer orderConn.Close()

    orderClient := orderpb.NewOrderServiceClient(orderConn)
    _, err = orderClient.CancelOrder(ctx, &orderpb.CancelOrderRequest{
        OrderId: req.OrderId,
    })
    if err != nil {
        log.Printf("Failed to cancel order for OrderID=%s: %v", req.OrderId, err)
        return &pb.CancelPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }
    log.Printf("Order cancelled successfully for OrderID=%s", req.OrderId)

    return &pb.CancelPaymentResponse{
        OrderId: req.OrderId,
        Status:  "CANCELLED",
    }, nil
}

// RefundPayment handles payment refunds
func (s *PaymentServiceServer) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the payment exists
    status, exists := s.paymentData[req.OrderId]
    if !exists || status != "COMPLETED" {
        log.Printf("Refund failed: OrderID=%s not found or not in COMPLETED state", req.OrderId)
        return &pb.RefundPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }

    // Update payment status to REFUNDED
    s.paymentData[req.OrderId] = "REFUNDED"
    log.Printf("Payment refunded: OrderID=%s, Status=REFUNDED", req.OrderId)

    return &pb.RefundPaymentResponse{
        OrderId: req.OrderId,
        Status:  "REFUNDED",
    }, nil
}

// SuccessPayment marks a payment as successful
func (s *PaymentServiceServer) SuccessPayment(ctx context.Context, req *pb.SuccessPaymentRequest) (*pb.SuccessPaymentResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if the payment exists
    status, exists := s.paymentData[req.OrderId]
    if !exists || status != "PENDING" {
        log.Printf("Success failed: OrderID=%s not found or not in PENDING state", req.OrderId)
        return &pb.SuccessPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }

    // Update payment status to COMPLETED
    s.paymentData[req.OrderId] = "COMPLETED"
    log.Printf("Payment successful: OrderID=%s, Status=COMPLETED", req.OrderId)

    // Call StartShipping in ShippingService
    shippingConn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Printf("Failed to connect to ShippingService: %v", err)
        return &pb.SuccessPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }
    defer shippingConn.Close()

    shippingClient := shippingpb.NewShippingServiceClient(shippingConn)
    _, err = shippingClient.StartShipping(ctx, &shippingpb.StartShippingRequest{
        OrderId: req.OrderId,
        Address: "123 Main St", // Replace with actual address logic
    })
    if err != nil {
        log.Printf("Failed to start shipping for OrderID=%s: %v", req.OrderId, err)
        return &pb.SuccessPaymentResponse{
            OrderId: req.OrderId,
            Status:  "FAILED",
        }, nil
    }
    log.Printf("Shipping started successfully for OrderID=%s", req.OrderId)

    return &pb.SuccessPaymentResponse{
        OrderId: req.OrderId,
        Status:  "COMPLETED",
    }, nil
}

func main() {
    // Create a listener on TCP port 50052
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the PaymentServiceServer
    paymentService := NewPaymentServiceServer()
    pb.RegisterPaymentServiceServer(grpcServer, paymentService)

    log.Println("PaymentService server is running on port 50052...")
    // Start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
