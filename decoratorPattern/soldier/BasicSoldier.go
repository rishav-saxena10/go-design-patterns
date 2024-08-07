package soldier

import "fmt"

type InterfaceSoldier interface {
	Attack() int
	Defence() int
}

type BasicSoldier struct{}

func (b BasicSoldier) Attack() int {
	return 1
}

func (b BasicSoldier) Defence() int {
	return 1
}

func DisplayStats(soldier InterfaceSoldier) {
	fmt.Printf("Soldier stats: Attack %d Defence %d", soldier.Attack(), soldier.Defence())
}
