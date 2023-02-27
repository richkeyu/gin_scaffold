package proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/richkeyu/gocommons/client"
	"github.com/richkeyu/gocommons/config"
	"github.com/richkeyu/gocommons/middleware"
	"github.com/richkeyu/gocommons/plog"
	"github.com/richkeyu/gocommons/server"

	"github.com/gin-gonic/gin"
)

type Proxy struct {
	config client.Config
	client *httputil.ReverseProxy
}

func (p *Proxy) Handler(method string, url string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		curUrl := url
		// 带参数的路由
		if len(ctx.Params) > 0 {
			for _, param := range ctx.Params {
				curUrl = strings.Replace(curUrl, ":"+param.Key, param.Value, 1)
			}
		}
		plog.Infof(server.NewContext(context.Background(), ctx), "gateway proxy %s - %s", ctx.Request.Host+ctx.Request.RequestURI, p.config.BaseUri+curUrl)
		ctx.Request.Header.Set("x-caller", "gateway")
		ctx.Request.Header.Set(middleware.GatewayHeaderKey, middleware.GatewaySecretKey) // 此值会用于后端服务访问来源判断
		ctx.Request.Method = method
		ctx.Request.URL.Path = curUrl
		p.client.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func (p *Proxy) SimpleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p.client.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func InitServiceByName(name string) (p *Proxy, err error) {
	var serviceConfigs map[string]client.Config
	err = config.Load(client.ServiceConfigName, &serviceConfigs)
	if err != nil {
		return nil, err
	}
	if c, ok := serviceConfigs[name]; ok {
		p = &Proxy{
			config: c,
			client: &httputil.ReverseProxy{
				Director: func(req *http.Request) {
					configUri, err := req.URL.Parse(c.BaseUri)
					if err != nil {
						return
					}
					req.URL.Scheme = configUri.Scheme
					req.URL.Host = configUri.Host
					req.Host = configUri.Host
				},
			},
		}
	} else {
		return nil, fmt.Errorf("not found %s service config", name)
	}
	return p, nil
}
