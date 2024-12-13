package main

import (
	"bytes"
	"database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		c.String(501, "This route is under development: %s\nDescription: %s", c.Request.URL, description)
	}
}

// Redirect to LinkedIn's auth front-end
func (app App) login(c *gin.Context) {
	state := 1
	c.Redirect(200, fmt.Sprintf(platform.LOGIN_URL_PATTERN, state, app.id, url.QueryEscape(app.authURL))) // Won't work immediatelly
}

// Register user and all 1st degree connections
func (app App) authPlatform(c *gin.Context) {
	loginToken, ok := c.Params.Get("code")
	if !ok {
		c.String(409, "Failed LinkedIn authorization. Missing auth code.")
	}

	_, ok = c.Params.Get("state")
	if !ok {
		c.String(500, "Failed LinkedIn authorization. Missing state check.")
	}

	// TODO: check the state in some sort of cache to see if it's in been called by our application.

	oAuthBody := map[string]string{
		"grant_type":    "authorization_code",
		"code":          loginToken,
		"redirect_uri":  url.QueryEscape(app.authURL),
		"client_id":     app.id,
		"client_secret": app.secret,
	}

	payload, err := json.Marshal(oAuthBody)
	if err != nil {
		c.String(500, "Failed to compose OAuth payload.")
	}

	res, err := http.Post(platform.OAUTH_URL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		c.String(500, "OAuth failed.")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.String(500, "Failed to read response body.")
	}

	var response struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		c.String(500, "Failed to parse response body.")
	}

	req, err := http.NewRequest(http.MethodGet, platform.PROFILE_SELF_URL, http.NoBody)
	if err != nil {
		c.String(500, "Failed to initialize request to profile endpoint: %s",  err.Error())
	}

	req.Header.Set("Authorization", "Bearer: " + response.AccessToken)




	// read code param
	// check if user is in platform
	// if it is, send some sort of credential
	// else process their data and their connections' and send credential

	// app.db.InsertAccount(&platform.Account{Username: "", PlatformId: ""})
}
