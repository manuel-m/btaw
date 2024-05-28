package main

import (
	"btaw/cmd/api/app"
	"btaw/cmd/api/cfg"
	"btaw/log"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app.Init()
	fmt.Println(cfg.Env)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	log.Logger.Fatal(err)

}
