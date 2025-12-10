package handlers

import (
	"api.teklifYonetimi/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusOK, "pong", gin.H{
		"message": "Server is running",
		"status":  "healthy",
	})
}

func (h *PingHandler) Hello(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusOK, "Hello endpoint", gin.H{
		"message": "Merhaba Go!",
	})
}
