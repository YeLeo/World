package poker

type Player struct {
	Name  string
	Cards []Card
}

func (p Player) CountCards() int {
	return len(p.Cards)
}

