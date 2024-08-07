package main

import (
	"fmt"
	"learning/go/decoratorPattern/soldier"
)

func main() {
	fmt.Println("Hello Decorator!!")

	basicSoldier := soldier.BasicSoldier{}
	soldier.DisplayStats(basicSoldier)
	fmt.Println()

	soldierWithSword := soldier.SoldierWithSword{
		Soldier: basicSoldier,
	}
	soldier.DisplayStats(soldierWithSword)
	fmt.Println()

	soldierWithShield := soldier.SoldierWithShield{
		Soldier: basicSoldier,
	}
	soldier.DisplayStats(soldierWithShield)
	fmt.Println()

	soldierWithSwordAndShield := soldier.SoldierWithShield{
		Soldier: soldier.SoldierWithSword{
			Soldier: basicSoldier,
		},
	}
	soldier.DisplayStats(soldierWithSwordAndShield)
	fmt.Println()

	fmt.Println("Bye Decorator!!")
}
