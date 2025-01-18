package main

import (
	"log"
	"math"
	"net/http"

	"github.com/davidbasilefilho/calculator-api.go/utils"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func createHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /add", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 { return f1 + f2 }))
	router.HandleFunc("GET /subtract", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 { return f1 - f2 }))
	router.HandleFunc("GET /multiply", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 { return f1 * f2 }))

	router.HandleFunc("GET /divide", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 {
		if f2 == 0 {
			log.Printf("cannot divide by zero")
			http.Error(w, "cannot divide by zero", http.StatusBadRequest)
		}
		return f1 / f2
	}))

	router.HandleFunc("GET /power", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 { return math.Pow(f1, f2) }))

	router.HandleFunc("GET /root", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 {
		if f2 == 0 {
			log.Printf("cannot divide by zero")
			http.Error(w, "cannot divide by zero", http.StatusBadRequest)
		}

		return math.Pow(f1, 1.0/f2)
	}))
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	createHandlers(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(router),
	}

	log.Printf("Server has started at %s", s.addr)
	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
