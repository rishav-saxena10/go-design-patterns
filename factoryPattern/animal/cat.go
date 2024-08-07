package animal

import "fmt"

type Cat struct {
	name string
}

func NewCat() *Cat {
	return &Cat{
		name: "Tom",
	}
}

func (c *Cat) Sound() {
	fmt.Printf("%s cat is making sound as Meow", c.name)
	fmt.Println()
}
