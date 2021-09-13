package main

type Customer struct {
	Email   string `json:"title"`
	PhoneNr string `json:"phonenr"`
	Views   int    `json:"views"`
}

type Event struct {
	PageHits   int `json:"pagehits"`
	VideoPlays int `json:"videoplays"`
}