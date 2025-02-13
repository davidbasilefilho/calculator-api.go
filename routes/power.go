package routes

import (
	"github.com/davidbasilefilho/calculator-api.go/utils"
	"math"
	"net/http"
)

func PowerHandler(router *http.ServeMux) {
	router.HandleFunc("GET /power", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 { return math.Pow(f1, f2) }))
}
