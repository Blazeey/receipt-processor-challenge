package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	host     string
	port     uint
	router   *gin.Engine
	shutDown chan struct{}
}

// Creating a new server and initializing the routes
func NewServer(host string, port uint) *Server {
	s := Server{
		host: host,
		port: port,
	}
	s.router = InitializeHttpRoutes()

	s.shutDown = make(chan struct{})
	return &s
}

// Starts the server
func (s *Server) Start() {
	fmt.Printf("Initializing server on port %d complete...\n", s.port)
	s.router.Run(fmt.Sprintf("%s:%d", s.host, s.port))
}

// Waits for shutdown
func (s *Server) WaitForShutdown() {
	<-s.shutDown
}
