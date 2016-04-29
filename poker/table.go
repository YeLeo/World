package poker

type Table struct {
	Players []*Player
	Deck    Deck
	Current Status
	History []*Status
	Banker  *Player
}

type Status struct {
	Cards  []*Card
	Player *Player
}

func NewTable() Table {
	t := Table{
		Players: make([]*Player, 0, 4),
		Deck:    NewDeck(),
		History: make([]*Status, 0)}
	return t
}

func (t *Table) AddPlayer(p *Player) {
	if t.Banker == nil {
		t.Banker = p
	}
	t.Players = append(t.Players, p)
}
