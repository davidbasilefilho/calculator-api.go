package routes

import (
	"github.com/davidbasilefilho/calculator-api.go/utils"
	"net/http"
)

func AddHandler(router *http.ServeMux) {
	router.HandleFunc("GET /add", utils.CreateHandler(func(f1, f2 float64, w http.ResponseWriter) float64 { return f1 + f2 }))
}
