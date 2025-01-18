package routes

import (
	"log"
	"net/http"
	"strconv"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	x, xErr := strconv.ParseFloat(query.Get("x"), 64)
	y, yErr := strconv.ParseFloat(query.Get("y"), 64)

	if xErr != nil {
		log.Printf("Error parsing 'x': %v", xErr)
		http.Error(w, "Invalid 'x' parameter", http.StatusBadRequest)
		return
	}

	if yErr != nil {
		log.Printf("Error parsing 'y': %v", yErr)
		http.Error(w, "Invalid 'y' parameter", http.StatusBadRequest)
		return
	}

	z := x + y

	log.Printf("Result: %v", z)
}
