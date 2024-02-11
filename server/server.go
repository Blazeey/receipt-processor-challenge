package server

import (
	"fmt"
	"receipt-processor-challenge/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	receipt *handler.ReceiptHandler
}

type Server struct {
	host               string
	port               uint
	router             *gin.Engine
	handlerMiddlewares []MiddlewareFunc
	shutDown           chan struct{}
}

func NewServer(host string, port uint) *Server {
	s := Server{
		host: host,
		port: port,
	}
	s.router = InitializeHttpRoutes()
	s.handlerMiddlewares = GetMiddlewares()

	s.shutDown = make(chan struct{})
	return &s
}

func (s *Server) Start() {
	s.router.Run(fmt.Sprintf("%s:%d", s.host, s.port))
}

func (s *Server) WaitForShutdown() {
	<-s.shutDown
}
