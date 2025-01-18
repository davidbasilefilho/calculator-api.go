package utils

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

func ParseInputs(w *http.ResponseWriter, r *http.Request) (float64, float64, error) {
	strX := r.URL.Query().Get("x")
	strY := r.URL.Query().Get("y")

	x, xErr := strconv.ParseFloat(strX, 64)
	y, yErr := strconv.ParseFloat(strY, 64)

	if xErr != nil {
		log.Printf("Error parsing 'x': %v", xErr)
		return 0, 0, errors.New("Invalid value for 'x'")
	}

	if yErr != nil {
		log.Printf("Error parsing 'y': %v", yErr)
		return 0, 0, errors.New("Invalid value for 'y'")
	}

	return x, y, nil
}
