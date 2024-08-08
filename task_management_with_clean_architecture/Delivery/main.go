package main

import (
	"context"
	"log"
	"os"
	"task_with_clean_arc/Delivery/router"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.TODO())
	router.CreateRouting(client)
}