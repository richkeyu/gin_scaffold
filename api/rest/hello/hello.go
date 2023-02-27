package hello

import (
	"context"
	"github.com/richkeyu/gocommons/plog"
	"github.com/richkeyu/gocommons/server"
	"github.com/richkeyu/gocommons/util/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 问候
// @Router / [GET]
func Greeter(c *gin.Context) {
	name, found := c.Get("name")
	if !found {
		name = "world"
	}

	ctx := server.NewContext(context.Background(), c)

	plog.Info(ctx, "Greeter")
	data := "hello " + name.(string)

	app.Response(c, http.StatusOK, 1, data)
}
