package api

import (
	"github.com/gin-gonic/gin"
)

func InitHandlers(r *gin.Engine) {
	r.GET("/status", Status)
}

func Status(c *gin.Context) {
	c.JSON(200, gin.H{"name": "hackaton-mephi-2024_backend", "status": "working"})
}
