package main

import (
	"git.wndv.co/thaneat.s/golang-di/example"
	"log"
)

func main() {
	connection, err := example.NewConnection()
	if err != nil {
		log.Println("create connection error")
	}

	orderRepository, err := example.NewOrderRepository(connection)
	if err != nil {
		log.Println("create order repository error")
	}

	foodService, err := example.NewFoodService(orderRepository)
	if err != nil {
		log.Println("create food service error")
	}

	app := example.Application{Name: "application", FoodService: foodService}
	if err != nil {
		log.Println("create application error")
	}

	log.Println("success...")
	log.Println(app.Name)
}
