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
	fmt.Println("1:WordCount 2:Shuffle 3:PlayCard")
	var t int
	fmt.Scan(&t)
	start := time.Now()
	switch t {
	case 1:
		wordcount.Do("C:/Users/M/Desktop/BigText/", "C:/Users/M/Desktop/BigText/Result-go.txt")
	case 2:
		deck := poker.NewDeck()
		m := make(map[poker.Card][]int)
		for i := 0; i < 100000; i++ {
			deck.Shuffle()
			for index, card := range deck.Cards {
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
	case 3:
		table := poker.NewTable()
		fmt.Println("请输入玩家1用户名：")
		var name string
		fmt.Scan(&name)
		table.AddPlayer(poker.Player{Name: name})
		for key, value := range table.Players {
			fmt.Println(key)
			fmt.Println(value.Name)
		}

	default:

	}
	fmt.Println("TotalTime:" + time.Since(start).String())
}
