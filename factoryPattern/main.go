package main

import (
	"fmt"
	"learning/go/factoryPattern/animal"
	"os"
)

func main() {
	fmt.Println("Hello Factory!!")

	input := "dog"

	var a animal.Animal

	switch input {
	case "cat":
		a = animal.NewCat()
	case "dog":
		a = animal.NewDog()
	default:
		fmt.Println("Invalid Animal!")
		os.Exit(1)
	}

	a.Sound()

	fmt.Println("Bye Factory!!")
}
