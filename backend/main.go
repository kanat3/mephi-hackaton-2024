package main

import (
	"backend/api"
	"backend/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.GetConfig()
	_ = cfg

	// init router
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r := gin.Default()
	r.Use(cors.New(config))
	api.InitHandlers(r)

	// run server
	r.Run(":8081")
}
