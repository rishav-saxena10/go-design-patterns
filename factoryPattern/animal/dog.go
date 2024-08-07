package animal

import "fmt"

type Dog struct {
	name string
}

func NewDog() *Dog {
	return &Dog{
		name: "Tuffi",
	}
}

func (c *Dog) Sound() {
	fmt.Printf("%s dog is making sound as Buff", c.name)
	fmt.Println()
}
