package poker

type Card struct {
	Value Rank
	Mark  Suit
}

func (c Card) ToString() string {
	return c.Mark.ToString() + c.Value.ToString()
}

type Rank int

const (
	ThreeRank Rank = iota
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
	TowRank
	BLACKJOKER
	REDJOKER
)

func (r Rank) ToString() string {
	switch r {
	case TowRank:
		return "2"
	case ThreeRank:
		return "3"
	case FourRank:
		return "4"
	case FiveRank:
		return "5"
	case SixRank:
		return "6"
	case SevenRank:
		return "7"
	case EightRank:
		return "8"
	case NineRank:
		return "9"
	case TenRank:
		return "10"
	case JACK:
		return "J"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	case ACE:
		return "A"
	case BLACKJOKER:
		return "joker"
	case REDJOKER:
		return "Joker"
	default:
		panic("Here can't find this rank!")
	}
}

type Suit int

const (
	JOKER Suit = iota
	SPADE
	HEART
	DIAMOND
	CLUB
)

func (s Suit) ToString() string {
	switch s {
	case JOKER:
		return ""
	case SPADE:
		return "♠"
	case HEART:
		return "♥"
	case DIAMOND:
		return "♦"
	case CLUB:
		return "♣"
	default:
		panic("Here can't find this suit!")
	}
}
