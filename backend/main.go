package main

import (
	"fmt"
	"os"
	"github.com/joakimsoren/screamdeal/backend/server"
)

func main() {
	server.StartServer()
	fmt.Println(fmt.Sprintf("The scariest API in the world runs on port %s, MOHAHAHAHAHAAHAH", os.Getenv("PORT")))
}
