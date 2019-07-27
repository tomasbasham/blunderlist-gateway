package main_test

//go:generate mockgen -destination=test/mock_blunderlist_todo_v1/todo.pb.go github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1 TodoClient,Todo_ListTasksClient
//go:generate mockgen -destination=test/mock_blunderlist_comment_v1/comment.pb.go github.com/tomasbasham/blunderlist-comment/blunderlist_comment_v1 CommentClient,Comment_ListCommentsClient

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"

	pb "github.com/golang/protobuf/ptypes"
	commentpbmock "github.com/tomasbasham/blunderlist-gateway/test/mock_blunderlist_comment_v1"
	todopbmock "github.com/tomasbasham/blunderlist-gateway/test/mock_blunderlist_todo_v1"
	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"

	"github.com/tomasbasham/blunderlist-gateway/internal/http"
	"github.com/tomasbasham/blunderlist-gateway/internal/store"
)

func TestProvider(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "blunderlist",
		Provider: "blunderlist-gateway",
	}

	// Create a new gomock controller. It defines the scope and lifetime of mock
	// objects, as well as their expectations.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock client types. The underlying store properties of these types will be
	// set for each pact intereaction.
	todoClient := todopbmock.NewMockTodoClient(ctrl)
	commentClient := commentpbmock.NewMockCommentClient(ctrl)

	// Start provider API in the background
	go startServer(todoClient, commentClient)

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		ProviderVersion: "0.0.0",

		// Pact broker credentials.
		BrokerURL:   os.Getenv("PACT_BROKER"),
		BrokerToken: os.Getenv("PACT_BROKER_TOKEN"),
		Tags:        strings.Split(os.Getenv("PACT_TAGS"), ","),

		// Push verification results to the pact broker.
		PublishVerificationResults: shouldPublishVerificationResults(),

		// Provider states.
		StateHandlers: types.StateHandlers{
			"a task exists": func() error {
				timestamp, err := pb.TimestampProto(time.Unix(656971200, 0))
				if err != nil {
					return err
				}

				todoStream := todopbmock.NewMockTodo_ListTasksClient(ctrl)
				todoStream.EXPECT().Recv().Return(&todopb.TaskResponse{
					Id:         1,
					Title:      "hello",
					Completed:  false,
					CreateTime: timestamp,
				}, nil)
				todoStream.EXPECT().Recv().Return(nil, io.EOF)

				commentStream := commentpbmock.NewMockComment_ListCommentsClient(ctrl)
				commentStream.EXPECT().Recv().Return(nil, io.EOF)

				any := gomock.Any()
				todoClient.EXPECT().ListTasks(any, any, any).Return(todoStream, nil)
				commentClient.EXPECT().ListComments(any, any, any).Return(commentStream, nil)

				return nil
			},
		},
	})
}

// startServer creates a minimal HTTP server with the same interface as the API
// under test.
func startServer(t *todopbmock.MockTodoClient, c *commentpbmock.MockCommentClient) {
	logger := log.New(ioutil.Discard, "", log.Lshortfile)
	store := store.New(t, c)

	gatewayService := http.NewServer(logger, store)
	if err := gatewayService.ServeHTTP("8080"); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func shouldPublishVerificationResults() bool {
	value, exists := os.LookupEnv("PACT_PUBLISH_VERIFICATION_RESULTS")
	if !exists {
		return false
	}

	b, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return b
}
