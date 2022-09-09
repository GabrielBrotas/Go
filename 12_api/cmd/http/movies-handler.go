package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

	// movie := models.Movie{
	// 	ID:          id,
	// 	Title:       "Some movie",
	// 	Description: "Best movie ever",
	// 	Year:        2020,
	// 	CreatedAt:   time.Date(2020, 55, 22, 01, 0, 0, 0, time.Local),
	// 	UpdatedAt:   time.Date(2020, 55, 22, 01, 0, 0, 0, time.Local),
	// }

	movie, err := app.models.DB.Get(id)

	if err != nil {
		app.logger.Print(errors.New(err.Error()))
		app.errorJSON(w, err)
		return
	}

	if movie == nil {
		app.errorJSON(w, errors.New("Movie not found"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

	if err != nil {
		app.logger.Print(errors.New(err.Error()))
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()

	if err != nil {
		app.errorJSON(w, err)
	}

	app.logger.Println(movies)

	err = app.writeJSON(w, http.StatusOK, movies, "movies")

	if err != nil {
		app.errorJSON(w, err)
	}
}

func (app *application) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie_input models.MovieInput

	err := json.NewDecoder(r.Body).Decode(&movie_input)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movie_id, err := app.models.DB.Create(movie_input)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, movie_id, "movie_id")
}
