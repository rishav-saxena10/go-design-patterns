package builder

import "fmt"

type Animal struct {
	name     string
	age      int64
	category string
}

func NewAnimal() *Animal {
	return &Animal{}
}

func NewAnimalWithFields(name string, age int64, category string) *Animal {
	return &Animal{
		name:     name,
		age:      age,
		category: category,
	}
}

func (a *Animal) WithName(name string) *Animal {
	a.name = name
	return a
}

func (a *Animal) WithAge(age int64) *Animal {
	a.age = age
	return a
}

func (a *Animal) WithCategory(category string) *Animal {
	a.category = category
	return a
}

func (a *Animal) PrintAnimal() {
	fmt.Printf("Animal: %+v", a)
	fmt.Println()
}
