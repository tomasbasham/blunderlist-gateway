package main

import (
	"context"
	"fmt"
	"log"
	"os"

	commentgrpc "github.com/tomasbasham/blunderlist-comment/grpc"
	todogrpc "github.com/tomasbasham/blunderlist-todo/grpc"

	"github.com/tomasbasham/blunderlist-gateway/internal/http"
	"github.com/tomasbasham/blunderlist-gateway/internal/store"
)

var todoServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("BLUNDERLIST_TODO_SERVICE_HOST"), os.Getenv("BLUNDERLIST_TODO_SERVICE_PORT"))
var commentServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("BLUNDERLIST_COMMENT_SERVICE_HOST"), os.Getenv("BLUNDERLIST_COMMENT_SERVICE_PORT"))

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile)

	// Create the client for the todo service using the address composed from the
	// the service's IP and port.
	todoClient, err := todogrpc.NewClientWithTarget(context.Background(), todoServiceAddr)
	if err != nil {
		logger.Fatalf("failed to connect: %v", err)
	}
	defer todoClient.Close()

	// Create the client for the comment service using the address composed from
	// the service's IP and port.
	commentClient, err := commentgrpc.NewClientWithTarget(context.Background(), commentServiceAddr)
	if err != nil {
		logger.Fatalf("failed to connect: %v", err)
	}
	defer commentClient.Close()

	store := store.New(todoClient, commentClient)

	gatewayService := http.NewServer(logger, store)
	gatewayServicePort := os.Getenv("BLUNDERLIST_GATEWAY_SERVICE_PORT")

	if err := gatewayService.ServeHTTP(gatewayServicePort); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
