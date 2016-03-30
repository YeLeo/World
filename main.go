package main

import (
	"fmt"
	"github.com/yeleo/world/wordcount"
	"time"
)

func main() {
	fmt.Println(time.Now())
	wordcount.Do()
	fmt.Println(time.Now())
}
