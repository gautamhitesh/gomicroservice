package main

import (
	. "awesomeProject/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World. Opening netflix...")
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movie", CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movie/{id}", GetMovie).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
