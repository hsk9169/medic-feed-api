package Routes

import (
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/medic-basic/s3-test/Controllers"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.Use(corsMiddleware())
	route.GET("/health/ready", Controllers.HealthCheck)

	route.POST("/feed/create", Controllers.PostFeed)
	route.GET("/feed/:patientId", Controllers.GetFeedList)
	route.GET("/feed/image/:patientId", Controllers.GetFeedImageList)

	route.POST("/auth/create", Controllers.PostAuthImage)
	route.GET("/auth/:medicId", Controllers.GetAuthImage)

	return route
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
