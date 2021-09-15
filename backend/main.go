package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/user", GetSingleUserById).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user", UpdateUser).Methods("PUT")
	router.HandleFunc("/user", DeleteUser).Methods("DELETE")

	cors := cors.New(cors.Options{})

	handler := cors.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))

}
