package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_PATTERN = "mongodb+srv://%s:%s@cluster0.3t4vm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment.")
	}

	connectionString := fmt.Sprintf(MONGODB_PATTERN, os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err = client.Database("development").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("MongoDB connection successful.")

	app := App{8000}
	app.Start()
}
