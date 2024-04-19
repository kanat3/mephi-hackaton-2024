package main

import (
	"backend/api"
	"backend/internal/config"

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
	r := gin.Default()
	api.InitHandlers(r)

	// run server
	r.Run(":8081")
}
