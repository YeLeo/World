package poker

type Card struct {
	Value Rank
	Mark  Suit
}

type Rank int

const (
	TowRank Rank = iota + 2
	ThreeRank
	FourRank
	FiveRank
	SixRank
	SevenRank
	EightRank
	NineRank
	TenRank
	JACK
	QUEEN
	KING
	ACE
	BLACKJOKER
	REDJOKER
)

type Suit int

const (
	JOKER Suit = iota
	SPADE
	HEART
	DIAMOND
	CLUB
)

func Deck() *[]Card {
	cards := make([]Card, 54)
	for rank := 2; rank < ACE; rank++ {
		for suit := 1; suit < CLUB; suit++ {
			append(cards, Card{Value: Rank(rank), Mark: Suit(suit)})
		}
	}
	append(cards, Card{Value: BLACKJOKER, Mark: JOKER})
	append(cards, Card{Value: REDJOKER, Mark: JOKER})
	return &cards
}
