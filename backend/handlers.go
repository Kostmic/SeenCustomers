package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	data, _ := json.Marshal(customers)
	writer.Write(data)
}
func GetSingleUserById(writer http.ResponseWriter, request *http.Request) {
	var userId UserId
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &userId)
	if err != nil {
		http.Error(writer, "Invalid id", http.StatusBadRequest) //400
		return
	}
	//Check id from request with customers in "DB"
	for _, customer := range customers {
		if customer.Id == userId.Id {
			response, _ := json.Marshal(customer)
			writer.Write(response)
			return
		}
	}
	//Return 402 no content if Id does not exist
	writer.WriteHeader(http.StatusNotFound) //404

}
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var customer Customer

	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &customer)
	if err != nil {
		http.Error(writer, "Invalid customer object", http.StatusBadRequest) //400
		return
	}

	//Generate unique id(UUID) for each customer
	customer.Id = uuid.New()
	customers = append(customers, customer)
	response, _ := json.Marshal(customer)
	writer.Write(response)

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var updCustomer Customer
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &updCustomer)

	if err != nil {
		http.Error(writer, "Invalid customer object", http.StatusBadRequest) //400
		return
	}

	log.Println(updCustomer.Id)
	//Check id from request with customers in "DB"
	for i := 0; i < len(customers); i++ {
		//Break out of loop if id is found
		if customers[i].Id == updCustomer.Id {
			customers[i].Email = updCustomer.Email
			customers[i].Phone = updCustomer.Phone
			customers[i].Views = updCustomer.Views
			response, _ := json.Marshal(updCustomer)
			writer.Write(response)
			return
		}
	}

	//If ID isn't found return 404
	http.Error(writer, "DENIED_ACCESS", http.StatusNotFound) //404
}
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var userId UserId
	
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &userId)
	if err != nil {
		http.Error(writer, "Invalid id", http.StatusBadRequest) //400
		return
	}
	
	//Check id from request with customers in "DB"
	for i := 0; i < len(customers); i++ {
		//Break out of loop if id is found
		if customers[i].Id == userId.Id {
			customers = append(customers[:i], customers[i+1:]...) // Truncate slice.
			return
		}
	}

	//If ID isn't found return 404
	http.Error(writer, "DENIED_ACCESS", http.StatusNotFound) //404
}
