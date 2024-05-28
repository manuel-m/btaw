package app

import (
	"btaw/cmd/api/cfg"
	"flag"
	"net/http"
)

var Mux *http.ServeMux

func Init() error {
	// cli parse
	{
		flag.IntVar(&cfg.Port, "port", 4000, "API server port")
		flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
		flag.Parse()
	}

	Mux = http.NewServeMux()

	return nil
}
