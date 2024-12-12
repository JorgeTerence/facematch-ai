package main

import (
	"database"
	"fmt"
	"net/url"
	"platform"

	"github.com/gin-gonic/gin"
)

type App struct {
	port    int `default:"8000"`
	db      database.DB
	id      string
	secret  string
	authURL string
}

func (app App) Start() {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/login", app.login)
		auth.POST("/connect", app.TODO("Respond to LinkedIn oauth redirect"))
		auth.POST("/logout", app.TODO("Log out from LinkedIn connection"))
		auth.POST("/update", app.TODO("Update your profile information"))
		auth.POST("/update-all", app.TODO("Update your network's profiles' information"))
	}

	service := r.Group("/match")
	{
		service.GET("/", app.TODO("Look for faces similar to yours in your LinkedIn network."))
		service.GET("/custom", app.TODO("Look for faces similar to a provided one in your LinkedIn network. Optionally include yourself"))
	}

	r.GET("/api", app.ping)
	r.Run(fmt.Sprintf(":%d", app.port))
}

// Check if the API is online
func (app App) ping(c *gin.Context) {
	c.Status(204)
}

// Fill-in handler for development purposes
func (app App) TODO(description string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.String(200, fmt.Sprintf("This route is under development: %s\nDescription: %s", c.Request.URL, description))
	}
}

// Redirect to LinkedIn's auth front-end
func (app App) login(c *gin.Context) {
	state := 1
	c.Redirect(200, fmt.Sprintf(platform.LOGIN_URL_PATTERN, state, app.id, url.QueryEscape(app.authURL))) // Won't work immediatelly
}

// Register user and all 1st degree connections
func (app App) authPlatform(c *gin.Context) {
	userToken, ok := c.Params.Get("code")
	if !ok {
		c.String(500, "Failed LinkedIn authorization. Missing auth code.")
	}

	state, ok := c.Params.Get("state")
	if !ok {
		c.String(500, "Failed LinkedIn authorization. Missing state check.")
	}

	// TODO: check the state in some sort of cache to see if it's in been called by our application.

	

	// read code param
	// check if user is in platform
	// if it is, send some sort of credential
	// else process their data and their connections' and send credential

	// app.db.InsertAccount(&platform.Account{Username: "", PlatformId: ""})
}
