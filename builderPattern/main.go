package main

import (
	"fmt"
	"learning/go/builderPattern/builder"
)

func main() {
	fmt.Println("Hello Builder !!")

	human1 := builder.NewHuman()

	human1.WithName("Rishav Saxena")
	human1.WithAge(27)
	human1.WithHeight(5.9)
	human1.WithWeight(56.5)

	human1.PrintHuman()

	human2 := builder.NewHumanWithFields("Devansh Gupta", 26, 5.8, 57.9)

	human2.PrintHuman()

	human3 := builder.NewHuman().WithName("Anant Agarwal").WithAge(27).WithHeight(6.1).WithWeight(89.7)

	human3.PrintHuman()

	animal1 := builder.NewAnimal()
	animal1.WithName("Cat")
	animal1.WithAge(27)
	animal1.WithCategory("Mammal")

	animal1.PrintAnimal()

	animal2 := builder.NewAnimalWithFields("Turtle", 10, "Amphibian")
	animal2.PrintAnimal()

	animal3 := builder.NewAnimal().WithName("Anaconda").WithAge(90).WithCategory("Reptile")
	animal3.PrintAnimal()

	fmt.Println("Bye Builder !!")
}
