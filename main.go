package main

import (
	"github.com/bsoviedo/go_backend_project/server/server"
)

func main() {

	server.App.Listen(":3000")

}
