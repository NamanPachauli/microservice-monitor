package main

type Service struct {
	ID     int    `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	URL    string `json:"url" bson:"url"`
	Status string `json:"status" bson:"status"`
}
