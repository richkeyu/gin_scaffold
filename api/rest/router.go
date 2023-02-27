package rest

import (
	"gateway/api/rest/hello"
	"gateway/pkg/middleware/header"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"github.com/richkeyu/gocommons/middleware"
)

var route gin.IRouter

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), middleware.InitTrace, middleware.LogReqResp, header.AddAcceptOriginHeader)

	r.GET("/", hello.Greeter)

	r.GET("/debug/stat/*filepath", func(context *gin.Context) {
		if context.Param("filepath") == "/ws" {
			statsviz.Ws(context.Writer, context.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/stat").ServeHTTP(context.Writer, context.Request)
	})

	route = r.Group("/api/v1/")

	registerRoute()

	return r
}
