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

	routes.AddHandler(router)
	routes.SubtractHandler(router)
	routes.MultiplyHandler(router)
	routes.DivideHandler(router)
	routes.PowerHandler(router)
	routes.RootHandler(router)

	chain := MiddlewareChain(RequestLoggerMiddleware)

	server := http.Server{
		Addr:    s.addr,
		Handler: chain(router),
	}

	log.Printf("Server has started at %s", s.addr)
	return server.ListenAndServe()
}
