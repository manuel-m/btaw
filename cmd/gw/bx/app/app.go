package app

import (
	"btaw"
	"btaw/pkg/exchange"
	"net/http"
)

const Version = btaw.Version

var Env = btaw.AppEnvDefault

var Port int
var DATABASE_URL string

var Exchange = &exchange.Bx{}
var Mux = http.NewServeMux()
