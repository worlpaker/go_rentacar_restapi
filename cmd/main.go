package main

import (
	"backend/api/server"
	"backend/config"
)

func main() {
	server.Start(config.Server_Port)
}
