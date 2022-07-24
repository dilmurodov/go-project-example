package api

import (
	"project/config"

	"project/api/handlers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "project/api/docs"
)

// @title           Swagger Example API
// SetUpRouter godoc
// @description     This is a simple server
// @termsOfService  http://swagger.io/terms/
// @host            localhost:5000
// @version         1.0
func SetUpRouter(h *handlers.Handler, cfg config.Config) (r *gin.Engine) {

	r = gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}