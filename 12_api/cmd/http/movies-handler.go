package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/GabrielBrotas/myapi/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("Id is ", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Some movie",
		Description: "Best movie ever",
		Year:        2020,
		CreatedAt:   time.Date(2020, 55, 22, 01, 0, 0, 0, time.Local),
		UpdatedAt:   time.Date(2020, 55, 22, 01, 0, 0, 0, time.Local),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

	if err != nil {
		app.logger.Print(errors.New(err.Error()))
	}

}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {}
