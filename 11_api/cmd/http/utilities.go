package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, key string) error {
	wrapper := make(map[string]interface{})

	wrapper[key] = data

	json_response, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json_response)

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	result_error := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, result_error, "error")
}
