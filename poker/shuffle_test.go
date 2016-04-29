package poker

import (
	"testing"
)

func BenchmarkShuffle(b *testing.B) {

	for i := 0; i < b.N; i++ {
		deck := NewDeck()
		deck.Shuffle()
		p := Player{Cards: deck.PopCards(13)}
		p.Name = "1"
		p.SortCards()
	}
	b.ReportAllocs()
}

func TestShuffle(t *testing.T) {
}
