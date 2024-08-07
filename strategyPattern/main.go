package main

import (
	"fmt"
	"learning/go/strategyPattern/strategy"
)

type Printer struct {
	strategy strategy.OutPutStrategy
}

func (p *Printer) setPrinter(strategy strategy.OutPutStrategy) {
	p.strategy = strategy
}

func main() {
	fmt.Println("Hello Strategy!!")

	s := "Coding Rocks!"

	printer := Printer{}

	printer.setPrinter(strategy.StringStrategy{})
	fmt.Println(printer.strategy.CreateOutput(s))

	printer.setPrinter(strategy.ByteStrategy{})
	fmt.Println(printer.strategy.CreateOutput(s))

	// printer.setPrinter(strategy.RuneStrategy{})
	// fmt.Println(printer.strategy.CreateOutput(s))

	fmt.Println("Bye Strategy!!")
}
