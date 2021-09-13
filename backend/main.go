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
 	router.HandleFunc("/user", CreateUser).Methods()
	router.HandleFunc("/user", UpdateUser).Methods()
	router.HandleFunc("/user", DeleteUser).Methods() 

	log.Fatal(http.ListenAndServe(":8080", router))

}