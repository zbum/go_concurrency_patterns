package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}

			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()

	return c
}

func cleanup() {
	fmt.Println("Cleaned up!!")
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)

}
