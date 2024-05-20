package main

import (
	"github.com/alionapermes/pf2sheet/internal/app/server"
)

func main() {
	serv := server.Server{}

	config := server.InitConfig()

	_ = serv.Start(config)
}
