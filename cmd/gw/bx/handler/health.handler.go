package handler

import (
	"btaw/cmd/gw/bx/app"
	"btaw/pkg"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Health(w http.ResponseWriter, r *http.Request) {

	jsonData := map[string]string{
		"status":      "available",
		"environment": app.Env,
		"version":     app.Version,
	}

	err := pkg.WriteJSON(w, http.StatusOK, jsonData, nil)

	if err != nil {
		log.Error().Err(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

}
