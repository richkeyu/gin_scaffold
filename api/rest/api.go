package rest

import (
	"gateway/api/rest/hello"
)

func registerRoute() {
	route.GET("/hello", hello.Greeter)
}
