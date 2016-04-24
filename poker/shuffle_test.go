package poker

import (
	"testing"
)

func BenchmarkShuffle(b *testing.B) {
	deck := NewDeck()
	for i := 0; i < b.N; i++ {
		deck.Shuffle()
	}
	b.ReportAllocs()
}

func TestShuffle(t *testing.T) {
}
