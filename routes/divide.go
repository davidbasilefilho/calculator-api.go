package routes

import (
	"github.com/davidbasilefilho/calculator-api.go/utils"
	"log"
	"net/http"
)

func DivideHandler(router *http.ServeMux) {
	router.HandleFunc("GET /divide", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 {
		if f2 == 0 {
			log.Printf("cannot divide by zero")
			http.Error(w, "cannot divide by zero", http.StatusBadRequest)
		}
		return f1 / f2
	}))
}
