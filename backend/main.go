package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: Start server
	fmt.Println(fmt.Sprintf("The scariest API in the world runs on port %s, MOHAHAHAHAHAAHAH", os.Getenv("PORT")))
}
