package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetSingleUserById).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
 	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE") 

	log.Fatal(http.ListenAndServe(":8080", router))

}