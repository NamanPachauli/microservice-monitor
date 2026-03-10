package main

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func startHealthChecker() {

	for {
		checkServices()
		time.Sleep(10 * time.Second)
	}

}

func checkServices() {

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return
	}

	var services []Service

	cursor.All(context.TODO(), &services)

	for _, service := range services {

		resp, err := http.Get(service.URL)

		status := "DOWN"

		if err == nil && resp.StatusCode == 200 {
			status = "UP"
		}

		collection.UpdateOne(
			context.TODO(),
			bson.M{"id": service.ID},
			bson.M{"$set": bson.M{"status": status}},
		)

	}

}
