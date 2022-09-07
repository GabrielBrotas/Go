package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := AppStatus{
		Version:     version,
		Environment: app.config.env,
		Status:      "Available",
	}

	json_response, err := json.Marshal(status)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)
}