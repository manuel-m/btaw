package main

import (
	"btaw"
	"btaw/cmd/gw/bx/app"
	"flag"
	"os"

	"btaw/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// .env parse
	{
		err := godotenv.Load()
		if err != nil {
			logger.Log.Fatal("Error loading .env file")
		}
		// [!] use os.LookupEnv instead
		app.DatabaseURL = os.Getenv("DatabaseURL")
	}

	// cli parse
	{
		flag.IntVar(&app.Port, "port", btaw.GwPortDefault, "API server port")
		flag.StringVar(&app.Env, "env", "development", "Environment (development|staging|production)")
		flag.Parse()
	}

	Routes()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      app.Mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Log.Printf("starting %s server on %s", app.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Log.Fatal(err)

}
