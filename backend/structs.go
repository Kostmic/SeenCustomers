package main

type Customer struct {
	Id      string `json:"id"`
	Email   string `json:"title"`
	PhoneNr string `json:"phonenr"`
	Views   int    `json:"views"`
}

type Event struct {
	PageHits   int `json:"pagehits"`
	VideoPlays int `json:"videoplays"`
}

type allCustomers []Customer

var customers = allCustomers{
	{
		Id:      "1",
		Email:   "john.eric@gmail.com",
		PhoneNr: "47391048",
		Views:   3,
	},
	{
		Id:      "2",
		Email:   "sarah_simpleton@hotmail.com",
		PhoneNr: "38593926",
		Views:   1,
	},
	{
		Id:      "3",
		Email:   "ChrisOlafson@gmail.com",
		PhoneNr: "19573628",
		Views:   8,
	},
}