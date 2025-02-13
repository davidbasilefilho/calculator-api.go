package utils

import (
	"fmt"
	"log"
	"net/http"
)

func CreateHandler(op func(float64, float64, http.ResponseWriter) float64) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		x, y, err := ParseInputs(&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := op(x, y, w)
		log.Printf("Result: %v", result)

		w.Write([]byte(fmt.Sprintf("%v", result)))
	}
}
