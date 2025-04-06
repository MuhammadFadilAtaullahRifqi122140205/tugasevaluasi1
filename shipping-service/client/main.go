package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/fadilrifqi/shipping-service/proto" // Import the generated proto package

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    // Connect to the ShippingService gRPC server
    conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to ShippingService: %v", err)
    }
    defer conn.Close()

    client := pb.NewShippingServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

    for {
        // Display the menu        fmt.Println("Choose an option:")
        fmt.Println("1. Mark Shipping as Shipped")
        fmt.Println("2. Cancel Shipping")
        fmt.Println("3. Exit")
        fmt.Print("Enter your choice: ")

        // Read user input
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            fmt.Println("Marking shipping as shipped...")
            handleShipped(client, "123") // Hardcoded Order ID
        case "2":
            fmt.Println("Cancelling shipping...")
            handleCancelShipping(client, "123") // Hardcoded Order ID
        case "3":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please select 1, 2, or 3.")
        }
    }
}

func handleShipped(client pb.ShippingServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call Shipped
    resp, err := client.Shipped(ctx, &pb.ShippedRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to mark shipping as shipped: %v", err)
        return
    }

    log.Printf("Shipped Response: OrderID=%s, Status=%s", resp.OrderId, resp.Status)
}

func handleCancelShipping(client pb.ShippingServiceClient, orderID string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call CancelShipping
    resp, err := client.CancelShipping(ctx, &pb.CancelShippingRequest{
        OrderId: orderID,
    })
    if err != nil {
        log.Printf("Failed to cancel shipping: %v", err)
        return
    }

    log.Printf("CancelShipping Response: Success=%v, Message=%s", resp.Success, resp.Message)
}
