package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func getServices(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var services []Service

	err = cursor.All(context.TODO(), &services)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(services)
}

func registerService(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var service Service

	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service.Status = "unknown"

	_, err = collection.InsertOne(context.TODO(), service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(service)
}
