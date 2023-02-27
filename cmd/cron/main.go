package main

import (
	"gateway/cmd/cron/cmd"
	_ "go.uber.org/automaxprocs"
)

func main() {
	//filePath := "config/app.yaml"
	//if os.Getenv("APP_ENV") == "dev" {
	//	filePath = "config/app_dev.yaml"
	//}

	cmd.Execute()
}
