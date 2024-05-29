package handler

import (
	"btaw/logger"
	"btaw/pkg/klines"
	"net/http"
)

func Klines(w http.ResponseWriter, r *http.Request) {
	symbol := r.PathValue("symbol")
	interval := r.PathValue("interval")
	json_bytes, err := klines.Fetch(symbol, interval)

	logger.Log.Printf("GET klines/%s/%s", symbol, interval)

	if err != nil {
		logger.Log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_bytes)

}
