package router

import (
	"github.com/gin-gonic/gin"
	"log-client-go/handlers"
)

// SetupRouter initializes the Gin router with routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	g := r.Group("/api")
	g.GET("/aggregations/user_agents", handlers.GetTopUserAgents)
	g.GET("/aggregations/http_methods", handlers.GetHTTPMethodsCount)
	g.GET("/aggregations/status_codes", handlers.GetStatusCodeDistribution)
	g.GET("/aggregations/device_names", handlers.GetDeviceNames)

	return r
}
