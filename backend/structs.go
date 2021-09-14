package main

import "github.com/google/uuid"

type Customer struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
	Views int       `json:"views"`
}
/* type Customers struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
	Views int       `json:"views"`
} */

type Event struct {
	PageHits   int `json:"pagehits"`
	VideoPlays int `json:"videoplays"`
}

type UserId struct {
	Id uuid.UUID `json:"id"`
}

type allCustomers []Customer

var customers = allCustomers{
	{
		Id:    uuid.New(),
		Email: "john.eric@gmail.com",
		Phone: "47391048",
		Views: 3,
	},
	{
		Id:    uuid.New(),
		Email: "sarah_simpleton@hotmail.com",
		Phone: "38593926",
		Views: 1,
	},
	{
		Id:    uuid.New(),
		Email: "ChrisOlafson@gmail.com",
		Phone: "19573628",
		Views: 8,
	},
}
