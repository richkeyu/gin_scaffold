package main

import (
	"fmt"
	"os"

	"gateway/api/rest"
	"gateway/pkg/db"
	"gateway/pkg/proxy"

	"github.com/richkeyu/gocommons/client"
	"github.com/richkeyu/gocommons/config"
	"github.com/richkeyu/gocommons/plog"
	"github.com/richkeyu/gocommons/server"
	"github.com/richkeyu/gocommons/trace"
	"github.com/richkeyu/gocommons/wrapper"
)

func main() {

	if _, err := config.NewConfig(os.Getenv(config.AppEnvName)); err != nil {
		panic(fmt.Sprintf("init config failed, err: %v", err))
	}

	plog.InitWithPath("log", "prod")

	client.AddDefaultWrappers(wrapper.HttpClientTrace)

	// 初始化Trace
	trace.InitGenerator()
	// 初始化代理
	proxy.InitProxy()
	// 初始化缓存
	db.InitLocalCache()

	// register route and init server
	router := rest.InitRouter()
	srv := server.NewServer(router)

	srv.Run()
}
