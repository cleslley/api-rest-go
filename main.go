package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cleslley/api-rest-go/config"
	"github.com/cleslley/api-rest-go/handler"
	movierouter "github.com/cleslley/api-rest-go/router"
	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("Server Starting...")
}

var data = handler.MoviesDB{}
var conf = config.Config{}

func init() {
	conf.Read()

	data.Server = conf.Mongo.Server
	fmt.Println("Server Name:", conf.Mongo.Server)
	data.Database = conf.Mongo.Database
	fmt.Println("Database:", conf.Mongo.Database)
	data.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", movierouter.GetAll).Methods("GET")
	r.HandleFunc("/movies/{id}", movierouter.GetByID).Methods("GET")
	r.HandleFunc("/movies", movierouter.Create).Methods("POST")
	r.HandleFunc("/movies/{id}", movierouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
