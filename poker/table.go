package poker

type Table struct {
	Players []Player
	Cards   Deck
	Current Status
	History []Status
}

type Status struct {
	Cards  []Card
	Player Player
}

func NewTable() Table {
	t := Table{
		Players: make([]Player, 0, 4),
		Cards:   NewDeck(),
		History: make([]Status, 0)}
	return t
}

func (t Table) AddPlayer(p Player) {
	//here likes should use a point?wait me test!
	t.Players = append(t.Players, p)
}
