package handlers

import (
	"project/api/http"
	"project/config"
	"project/pkg/logger"
	"project/storage"

	"github.com/gin-gonic/gin"
)

// Define Handler struct
type Handler struct {
	cfg  config.Config
	log  logger.LoggerI
	repo storage.StorageI
}
// Implement Handler Object 
func NewHandler(cfg config.Config, log logger.LoggerI, repo storage.StorageI) *Handler {
	return &Handler{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {

	// define log in different status codes 
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"!!!Response---->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("data", status.Description),
		)
	case code == 400:
		h.log.Warn(
			"!!!Response---->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response---->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("data", data),
		)
	}

	// JSON serialize the struct to response body and sets
	// COntent-type application/json 
	c.JSON(status.Code, http.Response{
		Status: status.Status,
		Description: status.Description,
		Data: data,
	})

}
