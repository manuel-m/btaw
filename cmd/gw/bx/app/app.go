package app

import (
	"btaw/pkg/exchange"
	"net/http"
)

const Version string = "0.0.1"

var Port int
var Env string
var DATABASE_URL string

var Exchange = &exchange.Bx{}
var Mux = http.NewServeMux()
