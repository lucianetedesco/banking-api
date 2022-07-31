package main

import (
	"github.com/lucianetedesco/banking-api/server"
	"github.com/lucianetedesco/banking-api/settings"
)

func main() {
	settings.InitConfig()

	server.CreateRoutesAndRun()
}
