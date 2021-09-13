package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)
func GetAllUsers() httprouter.Handle{
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		
	}
}
func GetSingleUserById() httprouter.Handle{
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		
	}
}
func CreateUser() httprouter.Handle{
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		
	}
}
func UpdateUser() httprouter.Handle{
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		
	}
}
func DeleteUser() httprouter.Handle{
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		
	}
}