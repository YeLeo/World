package poker

import (
	"math/rand"
	"time"
)

const deckSize int = 54

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	cards := make([]Card, 0, deckSize)
	for rank := TowRank; rank <= ACE; rank++ {
		for suit := SPADE; suit <= CLUB; suit++ {
			cards = append(cards, Card{Value: Rank(rank), Mark: Suit(suit)})
		}
	}
	cards = append(cards, Card{Value: BLACKJOKER, Mark: JOKER})
	cards = append(cards, Card{Value: REDJOKER, Mark: JOKER})
	return Deck{Cards: cards}
}

func (d Deck) Size() int {
	return len(d.Cards)
}

func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < d.Size(); i++ {
		index := r.Intn(d.Size()-i) + i
		d.Cards[i], d.Cards[index] = d.Cards[index], d.Cards[i]
	}
}

func (d Deck) PopCards(num int) []Card {
	cards := d.Cards[:num]
	d.Cards = d.Cards[num:]
	return cards
}
