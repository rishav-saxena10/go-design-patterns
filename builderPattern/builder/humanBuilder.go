package builder

import "fmt"

type Human struct {
	name   string
	age    int64
	height float64
	weight float64
}

func NewHuman() *Human {
	return &Human{}
}

func NewHumanWithFields(name string, age int64, height float64, weight float64) *Human {
	return &Human{
		name:   name,
		age:    age,
		height: height,
		weight: weight,
	}
}

func (h *Human) WithName(name string) *Human {
	h.name = name
	return h
}

func (h *Human) WithAge(age int64) *Human {
	h.age = age
	return h
}

func (h *Human) WithHeight(height float64) *Human {
	h.height = height
	return h
}

func (h *Human) WithWeight(weight float64) *Human {
	h.weight = weight
	return h
}

func (h *Human) PrintHuman() {
	fmt.Printf("Human: %+v", h)
	fmt.Println()
}
