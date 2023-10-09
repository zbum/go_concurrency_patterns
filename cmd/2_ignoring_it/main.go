package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boringWithRandom(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boringWithRandom("boring!")
}
