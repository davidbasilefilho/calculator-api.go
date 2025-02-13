package routes

import (
	"github.com/davidbasilefilho/calculator-api.go/utils"
	"log"
	"math"
	"net/http"
)

func RootHandler(router *http.ServeMux) {
	router.HandleFunc("GET /root", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 {
		if f2 == 0 {
			log.Printf("cannot divide by zero")
			http.Error(w, "cannot have an index of zero", http.StatusBadRequest)
		}

		return math.Pow(f1, 1.0/f2)
	}))
}
