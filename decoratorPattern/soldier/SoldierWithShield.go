package soldier

type SoldierWithShield struct {
	Soldier InterfaceSoldier
}

func (s SoldierWithShield) Attack() int {
	return s.Soldier.Attack() - 5
}

func (s SoldierWithShield) Defence() int {
	return s.Soldier.Defence() + 20
}
