package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllUsers(writer http.ResponseWriter, request *http.Request){
	json.NewEncoder(writer).Encode(customers)
}
func GetSingleUserById(writer http.ResponseWriter, request *http.Request){
	params := mux.Vars(request)
	//Check id from request with customers in "DB"
	for _, customer := range customers {
		if (customer.Id.String() == params["id"]) {
			json.NewEncoder(writer).Encode(customer)
			return
		}
	}
	//Return 402 no content if Id does not exist
	writer.WriteHeader(http.StatusNotFound) //404

}
func CreateUser(writer http.ResponseWriter, request *http.Request){
	var customer Customer

	err := json.NewDecoder(request.Body).Decode(&customer)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusBadRequest) //400
		return
	}
	//Generate unique id(UUID) for each customer
	customer.Id = uuid.New()

	customers = append(customers, customer)
	json.NewEncoder(writer).Encode(&customer)

}

func UpdateUser(writer http.ResponseWriter, request *http.Request){
	var updCustomer Customer
	params := mux.Vars(request)

	
	//Check id from request with customers in "DB"
	for i := 0; i < len(customers); i++ {
		//Break out of loop if id is found
		if (customers[i].Id.String() == params["id"]) {
			json.NewEncoder(writer).Encode(customers[i])
			err := json.NewDecoder(request.Body).Decode(&updCustomer)
			if err != nil{
				http.Error(writer, err.Error(), http.StatusBadRequest) //400
				return
			}
			customers[i].Email = updCustomer.Email
			customers[i].Phone = updCustomer.Phone
			customers[i].Views = updCustomer.Views
			return
		}
	}

	//If ID isn't found return 404
		http.Error(writer, "DENIED_ACCESS", http.StatusNotFound) //404
}
func DeleteUser(writer http.ResponseWriter, request *http.Request){
	params := mux.Vars(request)
	//Check id from request with customers in "DB"
	for i := 0; i < len(customers); i++ {
		//Break out of loop if id is found
		if (customers[i].Id.String() == params["id"]) {
			customers = append(customers[:i], customers[i+1:]...)   // Truncate slice.
			return
		}
	}

	//If ID isn't found return 404
	http.Error(writer, "DENIED_ACCESS", http.StatusNotFound) //404
}