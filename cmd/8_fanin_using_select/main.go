package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fan-in using select

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	//go func() {
	//	for {
	//		c <- <-input1
	//	}
	//}()
	//go func() {
	//	for {
	//		c <- <-input2
	//	}
	//}()

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func main() {

	c := fanIn(boring("Joe!"), boring("Ann!"))
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
