package header

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddAcceptOriginHeader 跨域
func AddAcceptOriginHeader(c *gin.Context) {
	if origin := c.GetHeader("Origin"); len(origin) > 0 {
		c.Header("Access-Control-Allow-Origin", origin)
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-requested-with")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
