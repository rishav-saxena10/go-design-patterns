package soldier

type SoldierWithSword struct {
	Soldier InterfaceSoldier
}

func (s SoldierWithSword) Attack() int {
	return s.Soldier.Attack() + 10
}

func (s SoldierWithSword) Defence() int {
	return s.Soldier.Defence()
}
