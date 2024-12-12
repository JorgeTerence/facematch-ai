package main

import (
	"log"
	"database"
	// TODO: "platform" package for connecting to the LinkedIn API

	"github.com/joho/godotenv"
)

// The app will need an actual privacy policy, because we're accessing their account data
// There'll be a front-end repo, not just some Go templates
// Both API Gateway and Amplify have free tiers. Mayyyyybe EC2 images are not ideal
// Perhaps I'll configure GitHub actions to handle CI/CD and deploy to AWS

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
