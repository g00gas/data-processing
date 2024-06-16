package router

import (
	"log-client-go/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router with routes
func SetupRouter(r *gin.Engine) *gin.Engine {
	g := r.Group("/api")
	g.GET("/aggregations/user_agents", handlers.GetTopUserAgents)
	g.GET("/aggregations/http_methods", handlers.GetHTTPMethodsCount)
	g.GET("/aggregations/status_codes", handlers.GetStatusCodeDistribution)
	g.GET("/aggregations/device_names", handlers.GetDeviceNames)

	return r
}
