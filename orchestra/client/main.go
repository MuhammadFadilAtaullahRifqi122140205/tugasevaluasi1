package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	orderpb "github.com/fadilrifqi/order-service/proto"       // Import the generated proto package for OrderService
	paymentpb "github.com/fadilrifqi/payment-service/proto"   // Import the generated proto package for PaymentService
	shippingpb "github.com/fadilrifqi/shipping-service/proto" // Import the generated proto package for ShippingService

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    // Connect to the OrderService gRPC server
    orderConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to OrderService: %v", err)
    }
    defer orderConn.Close()
    orderClient := orderpb.NewOrderServiceClient(orderConn)

    // Connect to the PaymentService gRPC server
    paymentConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to PaymentService: %v", err)
    }
    defer paymentConn.Close()
    paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)

    // Connect to the ShippingService gRPC server
    shippingConn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to ShippingService: %v", err)
    }
    defer shippingConn.Close()
    shippingClient := shippingpb.NewShippingServiceClient(shippingConn)

    reader := bufio.NewReader(os.Stdin)

    for {
        // Prompt for Order ID
        fmt.Print("Enter Order ID (or type 'exit' to quit): ")
        orderID, _ := reader.ReadString('\n')
        orderID = strings.TrimSpace(orderID)

        // Exit if the user types "exit"
        if strings.ToLower(orderID) == "exit" {
            fmt.Println("Exiting...")
            break
        }

        // Prompt for Address
        fmt.Print("Enter Address: ")
        address, _ := reader.ReadString('\n')
        address = strings.TrimSpace(address)

        // Call CreateOrder
        handleCreateOrder(orderClient, orderID)

        // Prompt for payment action
        fmt.Println("Choose an option:")
        fmt.Println("1. Cancel Payment")
        fmt.Println("2. Process Payment")
        fmt.Print("Enter your choice: ")
        paymentChoice, _ := reader.ReadString('\n')
        paymentChoice = strings.TrimSpace(paymentChoice)

        switch paymentChoice {
        case "1":
            handleCancelPayment(paymentClient, orderID)
            continue
        case "2":
            handleSuccessPayment(paymentClient, orderID)
        default:
            fmt.Println("Invalid choice. Returning to main menu.")
            continue
        }

        // Prompt for shipping action
        fmt.Println("Choose an option:")
        fmt.Println("1. Cancel Shipping")
        fmt.Println("2. Mark Shipping as Shipped")
        fmt.Print("Enter your choice: ")
        shippingChoice, _ := reader.ReadString('\n')
        shippingChoice = strings.TrimSpace(shippingChoice)

        switch shippingChoice {
        case "1":
            handleCancelShipping(shippingClient, orderID)
        case "2":
            handleShippedShipping(shippingClient, orderID)
        default:
            fmt.Println("Invalid choice. Returning to main menu.")
        }
    }
}

func handleCreateOrder(client orderpb.OrderServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call CreateOrder
    resp, err := client.CreateOrder(ctx, &orderpb.CreateOrderRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to create order: %v", err)
        return
    }

    log.Printf("CreateOrder Response: OrderID=%s, Status=%s", resp.OrderId, resp.Status)
}

func handleCancelPayment(client paymentpb.PaymentServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call CancelPayment
    resp, err := client.CancelPayment(ctx, &paymentpb.CancelPaymentRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to cancel payment: %v", err)
        return
    }

    log.Printf("CancelPayment Response: OrderID=%s, Status=%s", resp.OrderId, resp.Status)
}

func handleSuccessPayment(client paymentpb.PaymentServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call SuccessPayment
    resp, err := client.SuccessPayment(ctx, &paymentpb.SuccessPaymentRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to process payment: %v", err)
        return
    }

    log.Printf("SuccessPayment Response: OrderID=%s, Status=%s", resp.OrderId, resp.Status)
}

func handleCancelShipping(client shippingpb.ShippingServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call CancelShipping
    resp, err := client.CancelShipping(ctx, &shippingpb.CancelShippingRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to cancel shipping: %v", err)
        return
    }

    log.Printf("CancelShipping Response: Success=%v, Message=%s", resp.Success, resp.Message)
}

func handleShippedShipping(client shippingpb.ShippingServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call Shipped
    resp, err := client.Shipped(ctx, &shippingpb.ShippedRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to mark shipping as shipped: %v", err)
        return
    }

    log.Printf("Shipped Response: OrderID=%s, Status=%s", resp.OrderId, resp.Status)
}
