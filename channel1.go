package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// c := make(chan string)
	c :=  boring("boring!") // function returning a channel
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c ) // receive expression is just a value
	}
	fmt.Println("You're broing; I'm leaving.")
}

func boring(msg string) <-chan string {
        c := make(chan string)

	go func() { //  we launch the goroutine from inside the function 
	   for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // expression to be sent can be any suitable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	   }
	}()
        return c // Return the channel to the caller
}
