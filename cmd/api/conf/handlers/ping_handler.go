package handlers

import (
	"go-clean-arch/internal/platform/metrics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	metrics.ProcessBusiness.Inc()
	c.String(http.StatusOK, "pong")
}
