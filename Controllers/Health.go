package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medic-basic/s3-test/Services"
)

func HealthCheck(c *gin.Context) {
	fmt.Println("health check success")
	Services.Success(c, "ok", nil)
}
