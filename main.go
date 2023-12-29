package main

import (
	"fmt"
	"github.com/NestZ/golang-di/example"
)

func main() {
	app, cleanup, err := example.InitializeApplication("mongodb")
	if err != nil {
		fmt.Println(err)
	}
	defer cleanup()

	fmt.Println("success...")
	fmt.Println(app)
}
