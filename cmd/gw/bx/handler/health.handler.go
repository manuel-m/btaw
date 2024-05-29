package handler

import (
	"btaw/cmd/gw/bx/app"
	"btaw/logger"
	"btaw/pkg"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {

	json_data := map[string]string{
		"status":      "available",
		"environment": app.Env,
		"version":     app.Version,
	}

	err := pkg.WriteJSON(w, http.StatusOK, json_data, nil)

	if err != nil {
		logger.Log.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

}
