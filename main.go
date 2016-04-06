package main

import (
	"fmt"
	"github.com/yeleo/world/poker"
	"github.com/yeleo/world/wordcount"
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
		poker.Shuffle()
	default:
	}
	fmt.Println(time.Since(start))
	fmt.Scanln()
}
