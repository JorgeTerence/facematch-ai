package main

import (
	"log"
	"database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment.")
	}

	db := database.FromEnv()
	if err := db.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("MongoDB Connection Successful")

	app := App{8000}
	app.Start()
}
