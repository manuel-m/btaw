package app

import (
	"btaw/cmd/api/cfg"
	"btaw/logger"
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var Mux *http.ServeMux

func Init() error {

	// .env parse
	{
		err := godotenv.Load()
		if err != nil {
			logger.Log.Fatal("Error loading .env file")
		}
		// [!] use os.LookupEnv instead
		cfg.DATABASE_URL = os.Getenv("DATABASE_URL")
	}

	// cli parse
	{
		flag.IntVar(&cfg.Port, "port", 4000, "API server port")
		flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
		flag.Parse()
	}

	Mux = http.NewServeMux()

	return nil
}
