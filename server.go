package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	port int `default:"8000"`
}

func (app App) Start() {
	r := gin.Default()
	r.GET("/api", app.ping)
	r.Run(fmt.Sprintf(":%d", app.port))
}

func (app App) ping(c *gin.Context) {
	name := c.Query("name")

	if name != "" {
		c.String(200, fmt.Sprintf("Welcome, you kinda look like %s", name))
	} else {
		c.String(200, fmt.Sprintf("Welcome, you kinda look like <random actor from DB>"))
	}
}
