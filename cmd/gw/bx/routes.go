package main

import (
	"btaw/cmd/gw/bx/app"
	"btaw/cmd/gw/bx/handler"
)

func Routes() {
	app.Mux.HandleFunc("/health", handler.Health)
	app.Mux.HandleFunc("/klines/{symbol}/{tf}/{t0}/{duration}", handler.Klines)
}
