package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main(){
	router := httprouter.New()
	
	router.GET("/users", GetAllUsers())
	router.GET("/user/:id", GetSingleUserById())
	router.POST("/user", CreateUser())
	router.PUT("/user", UpdateUser())
	router.DELETE("/user", DeleteUser())
	log.Fatal(http.ListenAndServe(":8080", nil))
}