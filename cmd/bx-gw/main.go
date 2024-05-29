package main

import (
	"btaw/cmd/bx-gw/app"
	"btaw/cmd/bx-gw/cfg"
	"btaw/cmd/bx-gw/db"
	"btaw/cmd/bx-gw/handler"
	"btaw/logger"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app.Init()

	db.Query()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// routes
	{
		app.Mux.HandleFunc("/health", handler.Health)
		app.Mux.HandleFunc("/klines/{symbol}/{interval}", handler.Klines)
	}

	logger.Log.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Log.Fatal(err)

}