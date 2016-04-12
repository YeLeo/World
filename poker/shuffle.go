package poker

import (
	"math/rand"
	"time"
)

func Shuffle(cards *[]Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(*cards); i++ {
		index := r.Intn(len(*cards)-i) + i
		(*cards)[i], (*cards)[index] = (*cards)[index], (*cards)[i]
	}
}
