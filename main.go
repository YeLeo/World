package main

import (
	"fmt"
	"github.com/yeleo/world/wordcount"
	"time"
)

func main() {
	start := time.Now()
	wordcount.Do("C:/Users/M/Desktop/BigText/", "C:/Users/M/Desktop/BigText/Result-go.txt")
	fmt.Println(time.Since(start))
	fmt.Scanln()
}