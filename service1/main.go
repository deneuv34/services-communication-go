package main

import (
	"fmt"
	"service1/server"
)

func main() {
	fmt.Printf("Server is listening")
	server.StartServer()
}
