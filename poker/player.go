package poker

import (
	"fmt"
	"github.com/yeleo/world/util"
	"sort"
)

type Player struct {
	Name  string
	Cards []*Card
}

func (p Player) CountCards() int {
	return len(p.Cards)
}

func (p Player) ShowCards() {
	for _, card := range p.Cards {
		fmt.Println(card.ToString())
	}
}

func (p *Player) SortCards() {
	m := make(map[interface{}]interface{}, p.CountCards())
	for _, card := range p.Cards {
		m[card] = int(card.Value)
	}
	sm := util.CreateSM(m, util.SortByValue, util.ASC)
	sort.Sort(sm)
	for i := 0; i < len(sm.KeyValuePairs); i++ {
		p.Cards[i] = sm.KeyValuePairs[i].Key.(*Card)
	}
}
