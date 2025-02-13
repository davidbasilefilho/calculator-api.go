package main

import (
	"github.com/davidbasilefilho/calculator-api.go/routes"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	AddHandler(router)

	chain := MiddlewareChain(RequestLoggerMiddleware)

	server := http.Server{
		Addr:    s.addr,
		Handler: chain(router),
	}

	log.Printf("Server has started at %s", s.addr)
	return server.ListenAndServe()
}
