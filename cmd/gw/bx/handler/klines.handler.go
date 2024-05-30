package handler

import (
	"btaw/cmd/gw/bx/app"
	"btaw/logger"
	"btaw/pkg/timeutil"
	"fmt"
	"net/http"
	"strconv"
)

func Klines(w http.ResponseWriter, r *http.Request) {
	symbol := r.PathValue("symbol")
	tf := timeutil.Interval(r.PathValue("tf"))

	duration := timeutil.Interval(r.PathValue("duration"))
	// duration_ms, _ := duration.ToMs()
	// t0 := (time.Now().UnixNano() / 1e6) - duration_ms

	t0, err := strconv.ParseInt(r.PathValue("t0"), 10, 64)
	if err != nil {
		fmt.Println("Erreur de conversion:", err)
		http.Error(w, "invalid t0", http.StatusBadRequest)
		return
	}

	jsonBytes, err := app.Exchange.Klines(symbol, tf, t0, duration)
	logger.Log.Printf("GET klines/%s/%s/%d/%s", symbol, tf, t0, duration)

	if err != nil {
		logger.Log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
