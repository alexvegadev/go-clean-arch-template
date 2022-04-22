package conf

import (
	"go-clean-arch/cmd/api/conf/handlers"
	"go-clean-arch/internal/business"
	"go-clean-arch/internal/platform/log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(dep *dependencies) *gin.Engine {
	r := gin.New()

	//Middlewares
	r.Use(gin.LoggerWithWriter(log.GetOut(), "/ping", "/health"))

	//API group
	v1 := r.Group("/api/v1")

	//Use cases
	businessUseCase := business.NewBusinessUseCase(dep.DB)

	//Handlers
	businessHandler := handlers.NewBusinessHandler(businessUseCase)

	//Endpoints
	v1.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1.POST("/business", businessHandler.HandleCreate)

	v1.GET("/ping", handlers.Ping)

	return r
}
