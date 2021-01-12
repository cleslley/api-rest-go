package router

import (
	"encoding/json"
	"net/http"

	"github.com/cleslley/api-rest-go/handler"
	"github.com/cleslley/api-rest-go/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var data = handler.MoviesDB{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//GetAll retorna todos os dados
func GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := data.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, movies)
}

//GetByID retorna um dado por ID
func GetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movie, err := data.GetByID(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJSON(w, http.StatusOK, movie)
}

//Create Insere um novo dado
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := data.Create(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, movie)
}

//Delete remove um dado via id
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := data.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
