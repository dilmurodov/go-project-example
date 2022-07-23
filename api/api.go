package api

import (
	"project/config"

	"project/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(h *handlers.Handler, cfg config.Config) (r *gin.Engine) {
	
	r = gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)

	return
}