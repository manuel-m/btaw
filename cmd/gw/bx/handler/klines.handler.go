package handler

import (
	"btaw/cmd/gw/bx/app"
	"btaw/logger"
	"net/http"
)

func Klines(w http.ResponseWriter, r *http.Request) {
	symbol := r.PathValue("symbol")
	interval := r.PathValue("interval")
	json_bytes, err := app.Exchange.Klines(symbol, interval)
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
