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
				if m[*card] == nil {
					m[*card] = make([]int, 54)
				}
				m[*card][index]++
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
		for {
			if len(table.Players) < 4 {
				fmt.Printf("请输入玩家%d用户名：", len(table.Players)+1)
				var name string
				fmt.Scan(&name)
				table.AddPlayer(&poker.Player{Name: name})
			} else {
				break
			}
		}
		table.Deck.Shuffle()
		for _, player := range table.Players {
			player.Cards = table.Deck.PopCards(13)
		}
		table.Banker.Cards = append(table.Banker.Cards, table.Deck.Cards)
		for _, player := range table.Players {
			player.SortCards()
			fmt.Println(player.Name)
		}

	default:
		//Pop(13) is random?Yes,though i know,test it
		table := poker.NewTable()
		result := make(map[*poker.Player][]int, 4)
		for {
			if len(table.Players) < 4 {
				fmt.Printf("请输入玩家%d用户名：", len(table.Players)+1)
				var name string
				fmt.Scan(&name)
				player := &poker.Player{Name: name}
				table.AddPlayer(player)
				result[player] = make([]int, 15)
			} else {
				break
			}
		}

		for i := 0; i < 100000; i++ {
			table.Deck = poker.NewDeck()
			table.Deck.Shuffle()
			for player, record := range result {
				player.Cards = table.Deck.PopCards(13)
				for _, card := range player.Cards {
					record[int(card.Value)]++
				}
			}
		}
		for player, record := range result {
			fmt.Println(player.Name)
			fmt.Println(record)
		}

	}
	fmt.Println("TotalTime:" + time.Since(start).String())
}
