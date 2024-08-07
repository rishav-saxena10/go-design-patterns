package strategy

import "fmt"

type OutPutStrategy interface {
	CreateOutput(s string) string
}

type StringStrategy struct{}

func (s StringStrategy) CreateOutput(str string) string {
	fmt.Println("String Strategy")
	return str
}

type ByteStrategy struct{}

func (b ByteStrategy) CreateOutput(str string) string {
	fmt.Println("Byte Strategy")
	return fmt.Sprintf("%v", []byte(str))
}

type RuneStrategy struct{}

// func (r RuneStrategy) CreateOutput(str string) string {
// 	fmt.Println("Rune Strategy")
// 	return fmt.Sprintf("%v", []rune(str))
// }
