package main

import (
	"fmt"
	"git.wndv.co/thaneat.s/golang-di/example"
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
