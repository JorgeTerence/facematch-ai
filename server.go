package main

import (
	"database"
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	port int `default:"8000"`
	db   database.DB
}

func (app App) Start() {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/register", app.TODO("Register user and all 1st degree connections"))
		auth.POST("/login", app.TODO("Respond to LinkedIn oauth redirect"))
		auth.POST("/logout", app.TODO("Log out from LinkedIn connection"))
	}

	service := r.Group("/match")
	{
		service.GET("/", app.TODO("Look for similar faces in LinkedIn network."))
	}

	r.GET("/api", app.ping)
	r.Run(fmt.Sprintf(":%d", app.port))
}

// Check if the API is online
func (app App) ping(c *gin.Context) {
	name := c.Query("name")

	if name != "" {
		c.String(200, fmt.Sprintf("Welcome, you kinda look like %s", name))
	} else {
		c.String(200, fmt.Sprintf("Welcome, you kinda look like <random actor from DB>"))
	}
}

// Fill-in handler for development
func (app App) TODO(description string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.String(200, fmt.Sprintf("This route is under development: %s\nDescription: %s", c.Request.URL, description))
	}
}

func (app App) register(c *gin.Context) {
	app.db.InsertAccount(&database.Account{Username: "", PlatformId: 1})
}
