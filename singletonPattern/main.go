package main

import (
	"fmt"
	"learning/go/singletonPattern/singletonService"
)

func main() {
	fmt.Println("Hello Singleton!!")

	singletonIdService := singletonService.SingletonIdService{}

	idService1 := singletonIdService.NewSingletonIdService()

	buildCar(idService1.Next())
	buildBike(idService1.Next())

	idService2 := singletonIdService.NewSingletonIdService()

	buildCar(idService2.Next())
	buildBike(idService2.Next())

	buildCar(idService1.Next())

	fmt.Println("Bye Singleton!!")
}

func buildCar(id int) {
	fmt.Printf("Building Car with Id: %d", id)
	fmt.Println()
}

func buildBike(id int) {
	fmt.Printf("Building Bike with Id: %d", id)
	fmt.Println()
}
