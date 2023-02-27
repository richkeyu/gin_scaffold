package middleware

import (
	"gateway/pkg/constant"
	commonmiddleware "github.com/richkeyu/gocommons/middleware"
	"github.com/gin-gonic/gin"
	"strings"
)

func IpAuthMiddleware() gin.HandlerFunc {
	return commonmiddleware.IpAuthMiddleWare(constant.AllowIpList, func(c *gin.Context, isPass bool) (newIsPass bool) {
		// 如果ip不符合临时检查header是否有token
		if !isPass {
			auth := c.Request.Header.Get("authorization")
			if len(auth) > 0 && strings.HasPrefix(auth, "Bearer ") {
				isPass = true
			}
		}
		return isPass
	})
}
