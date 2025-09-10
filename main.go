package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"CRUD_operation/Controller"
	"CRUD_operation/router"
	"CRUD_operation/service"
	"CRUD_operation/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("CrudOperationGolnag")
	collection := db.Collection("users")

	utils.CleanupUsersCollection(ctx, collection)

	userService := service.NewUserService(collection)
	userHandler := Controller.NewUserHandler(userService)

	r := router.NewRouter(userHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
