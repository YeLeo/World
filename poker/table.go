package poker

type Table struct {
	Players []*Player
	Deck    Deck
	Current Status
	History []*Status
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
	t.Players = append(t.Players, p)
}

func (t *Table) NewRound() {

}
