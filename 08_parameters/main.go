package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port, numberOfRetries int) {
	println("Starting server...")
	// do important things
	println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}