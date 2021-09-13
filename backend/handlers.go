package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(writer http.ResponseWriter, request *http.Request){
	json.NewEncoder(writer).Encode(customers)
}
func GetSingleUserById(writer http.ResponseWriter, request *http.Request){
	params := mux.Vars(request)
	for _, customer := range customers {
		if customer.Id == params["id"] {
			json.NewEncoder(writer).Encode(customer)
		}
	}

}
func CreateUser(writer http.ResponseWriter, request *http.Request){

}

func UpdateUser(writer http.ResponseWriter, request *http.Request){

}
func DeleteUser(writer http.ResponseWriter, request *http.Request){

}