package main

import (
	"log"
	"os"
	"env"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment.")
	}

	db := DB{os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"), nil, env.FromString(os.Getenv("ENV"))}
	if err := db.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("MongoDB Connection Successful")

	app := App{8000}
	app.Start()
}
