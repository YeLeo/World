package main

import (
	"fmt"
	"github.com/yeleo/world/poker"
	"github.com/yeleo/world/wordcount"
	"strconv"
	"time"
)

func main() {
	fmt.Println("YeLeo's world!")
	fmt.Println("1:wordcount 2:Shuffle")
	var t int
	fmt.Scan(&t)
	start := time.Now()
	switch t {
	case 1:
		wordcount.Do("C:/Users/M/Desktop/BigText/", "C:/Users/M/Desktop/BigText/Result-go.txt")
	case 2:
		cards := poker.InitCards()
		m := make(map[poker.Card][]int)
		for i := 0; i < 100000; i++ {
			poker.Shuffle(cards)
			for index, card := range *cards {
				if m[card] == nil {
					m[card] = make([]int, 54)
				}
				m[card][index]++
			}
		}
		for key, value := range m {
			if key.Mark == poker.JOKER {
				fmt.Print(key.ToString() + "\t->")
			} else {
				fmt.Print(key.ToString() + "\t\t->")
			}
			for _, num := range value {
				fmt.Print(strconv.Itoa(num) + "\t")
			}
			fmt.Print("\n")
		}
	default:
	}
	fmt.Println("TotalTime:" + time.Since(start).String())
	//fmt.Scanln()
	//fmt.Scanln()
}
