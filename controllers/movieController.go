package controllers

import (
	. "awesomeProject/models"
	. "awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type MovieModel struct {
	*Movie
}

func (m *MovieModel) isEmpty() bool {
	return m.MovieName == ""
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Hello World")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one movie title")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty data fields")
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if movie.MovieName == "" {
		json.NewEncoder(w).Encode("JSON data empty")
		return
	}

	rand.Seed(time.Now().UnixNano())
	movie.UniqueId = strconv.Itoa(rand.Intn(100))
	CreateOneMovie(&movie)
	json.NewEncoder(w).Encode("Successfully added movie title")
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all movie titles")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(GetAllMovies())
	if err != nil {
		log.Fatal(err)
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Println("Getting movie details", movieId)
	err := json.NewEncoder(w).Encode(GetMovieDetails(movieId))
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty body")
		return
	}
	var movieDetails Movie
	_ = json.NewDecoder(r.Body).Decode(&movieDetails)
	if movieDetails.MovieName == "" {
		json.NewEncoder(w).Encode("Movie not found")
		return
	}
	fmt.Println("Updating movie details for", movieId, movieDetails)
	err := json.NewEncoder(w).Encode(UpdateMovieDetails(movieId, &movieDetails))
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Println("Deleting movie with id", movieId)
	err := json.NewEncoder(w).Encode(DeleteMovieRecord(movieId))
	if err != nil {
		log.Fatal(err)
	}
}
