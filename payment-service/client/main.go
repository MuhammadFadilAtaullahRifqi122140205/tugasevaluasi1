package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	orderpb "github.com/fadilrifqi/payment-service/proto"
	paymentpb "github.com/fadilrifqi/payment-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Failed() {
    // Connect to the PaymentService gRPC server
    conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to PaymentService: %v", err)
    }
    defer conn.Close()

    client := paymentpb.NewPaymentServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call CancelPayment
    cancelPaymentResp, err := client.CancelPayment(ctx, &paymentpb.CancelPaymentRequest{
        OrderId: "123",
    })
    if err != nil {
        log.Fatalf("Failed to cancel payment: %v", err)
    }

    log.Printf("CancelPayment Response: OrderID=%s, Status=%s", cancelPaymentResp.OrderId, cancelPaymentResp.Status)
}

func Success() {
    // Connect to the PaymentService gRPC server
    conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to PaymentService: %v", err)
    }
    defer conn.Close()

    client := orderpb.NewPaymentServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Call SuccessPayment
    successPaymentResp, err := client.SuccessPayment(ctx, &orderpb.SuccessPaymentRequest{
        OrderId: "123",
    })
    if err != nil {
        log.Fatalf("Failed to mark payment as successful: %v", err)
    }

    log.Printf("SuccessPayment Response: OrderID=%s, Status=%s", successPaymentResp.OrderId, successPaymentResp.Status)
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        // Display the menu
        fmt.Println("Choose an option:")
        fmt.Println("1. Failed Payment")
        fmt.Println("2. Success Payment")
        fmt.Println("3. Exit")
        fmt.Print("Enter your choice: ")

        // Read user input
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        // Trim and parse the input
        choice, err := strconv.Atoi(strings.TrimSpace(input))
        if err != nil {
            fmt.Println("Invalid input. Please enter a number.")
            continue
        }

        // Handle the user's choice
        switch choice {
        case 1:
            fmt.Println("Running Failed Payment...")
            Failed() // Call the failed payment function
        case 2:
            fmt.Println("Running Success Payment...")
            Success() // Call the success payment function
        case 3:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please select 1, 2, or 3.")
        }
    }
}
