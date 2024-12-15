package main

import (
	"fmt"
	"hexagonal-potato/logger"
	"hexagonal-potato/ports/input"
)

func main() {
	err := logger.Init("app.log")
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}
	input.InitHTTPService()
}
