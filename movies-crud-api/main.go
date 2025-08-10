package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// json: is a struct tag in Go used to tell the encoding/json package
// how to handle a struct field when encoding it into JSON or decoding JSON into the struct.
// It's not a part of the Go language itself but a convention followed by the standard library.

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// Set the Content-Type header to application/json and send the json encoded movie slice
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get a movie by id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Catch the parameters from the request
	params := mux.Vars(r)

	// Loop over existing movies
	for _, item := range movies {
		// If ID matches a the ID of an existing movie
		if item.ID == params["id"] {
			// JSON encode the item and send it
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "No movie found with the given ID", http.StatusNotFound)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// Delete a movie by id
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Catch the parameters from the request
	params := mux.Vars(r)

	// Loop over existing movies
	for index, item := range movies {
		// If ID matches a the ID of an existing movie
		if item.ID == params["id"] {
			// Delete it from the slice by appending the first items till index, escape index, append the rest
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// JSON encode the movie slice and send it
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	// Existing demo data in movie slice
	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "12345678",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "87654321",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Jane",
			Lastname:  "Doe",
		},
	})

	// Create methods
	r.HandleFunc("/movies", createMovie).Methods("POST")

	// Read methods
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")

	// Update methods
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	// Delete methods
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	// Staring server
	fmt.Println("Starting server at port 8000")

	// Log if server doesn't start
	log.Fatal(http.ListenAndServe(":8000", r))
}
