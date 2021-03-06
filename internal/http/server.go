package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tomasbasham/blunderlist-gateway/internal/store"
	httptransport "github.com/tomasbasham/grpc-service-go/transport/http"
)

// Server is a type composed of multiple endpoints and their respective
// handlers to serve HTTP requests.
type Server struct {
	logger *log.Logger
	store  *store.Store
}

// NewServer creates a new Server.
func NewServer(logger *log.Logger, store *store.Store) *Server {
	return &Server{
		logger: logger,
		store:  store,
	}
}

// ServeHTTP listens on a specific port across all TCP network interfaces.
func (gs *Server) ServeHTTP(port string) error {
	server := httptransport.NewServer(gs.createRouter())
	server.Addr = fmt.Sprintf(":%s", port)

	gs.logger.Printf("server started on [::]:%s", port)
	return server.ListenAndServe()
}

func (gs *Server) createRouter() http.Handler {
	router := mux.NewRouter()

	// Create a new middleware stack that sets the HTTP Accept and Content-Type
	// headers.
	middleware := httptransport.Chain(withAccept, withContentType)

	tasks := router.PathPrefix("/tasks").Subrouter()
	comments := router.PathPrefix("/comments").Subrouter()

	// Tasks
	tasks.HandleFunc("", middleware(gs.listTasks)).Methods("GET")
	tasks.HandleFunc("/{id:[0-9]+}", middleware(gs.getTask)).Methods("GET")
	tasks.HandleFunc("", middleware(gs.createTask)).Methods("POST")
	tasks.HandleFunc("/{id:[0-9]+}", middleware(gs.updateTask)).Methods("PUT", "PATCH")
	tasks.HandleFunc("/{id:[0-9]+}", middleware(gs.deleteTask)).Methods("DELETE")
	tasks.HandleFunc("/{id:[0-9]+}/comments", middleware(gs.listComments)).Methods("GET")

	// Comments
	comments.HandleFunc("/{id:[0-9]+}", middleware(gs.getComment)).Methods("GET")
	comments.HandleFunc("", middleware(gs.createComment)).Methods("POST")
	comments.HandleFunc("/{id:[0-9]+}", middleware(gs.updateComment)).Methods("PUT", "PATCH")
	comments.HandleFunc("/{id:[0-9]+}", middleware(gs.deleteComment)).Methods("DELETE")

	// Health checking
	router.HandleFunc("/healthz", gs.healthCheck).Methods("GET")

	return router
}
