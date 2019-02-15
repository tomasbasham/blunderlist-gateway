package main

import (
	"context"
	"fmt"
	"log"
	"os"

	commentgrpc "github.com/tomasbasham/blunderlist-comment/grpc"
	todogrpc "github.com/tomasbasham/blunderlist-todo/grpc"

	"github.com/tomasbasham/blunderlist-gateway/internal/comment"
	"github.com/tomasbasham/blunderlist-gateway/internal/http"
	"github.com/tomasbasham/blunderlist-gateway/internal/store"
	"github.com/tomasbasham/blunderlist-gateway/internal/todo"
)

var todoServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("BLUNDERLIST_TODO_SERVICE_HOST"), os.Getenv("BLUNDERLIST_TODO_SERVICE_PORT"))
var commentServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("BLUNDERLIST_COMMENT_SERVICE_HOST"), os.Getenv("BLUNDERLIST_COMMENT_SERVICE_PORT"))

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile)

	// Create the client for the todo service using the address composed from
	// the the service's IP and port.
	todoService, err := todogrpc.NewClientWithTarget(context.Background(), todoServiceAddr)
	if err != nil {
		logger.Fatalf("failed to connect: %v", err)
	}
	defer todoService.Close()

	// Create the client for the comment service using the address composed
	// from the service's IP and port.
	commentService, err := commentgrpc.NewClientWithTarget(context.Background(), commentServiceAddr)
	if err != nil {
		logger.Fatalf("failed to connect: %v", err)
	}
	defer commentService.Close()

	// Create the internal client types that embed the gRPC clients and
	// encapsulate the conversion between the gRPC wire format and internal
	// application types.
	todoClient := todo.NewClient(logger, todoService)
	commentClient := comment.NewClient(logger, commentService)

	// Create store types that encapsulate the application logic.
	config := http.ServerConfig{
		CreateCommentStore: store.NewCreateComment(commentClient),
		CreateTaskStore:    store.NewCreateTask(todoClient),
		DeleteCommentStore: store.NewDeleteComment(commentClient),
		DeleteTaskStore:    store.NewDeleteTask(todoClient),
		GetCommentStore:    store.NewGetComment(commentClient),
		GetTaskStore:       store.NewGetTask(todoClient, commentClient),
		ListCommentsStore:  store.NewListComments(commentClient),
		ListTasksStore:     store.NewListTasks(todoClient, commentClient),
		UpdateCommentStore: store.NewUpdateComment(commentClient),
		UpdateTaskStore:    store.NewUpdateTask(todoClient),
	}

	gatewayService := http.NewServer(logger, config)
	gatewayServicePort := os.Getenv("BLUNDERLIST_GATEWAY_SERVICE_PORT")

	if err := gatewayService.ServeHTTP(gatewayServicePort); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
