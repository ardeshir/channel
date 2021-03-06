package main

import (
	"fmt"
	"time"
	"math/rand"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
        
	return c
}


func main() {
	/* c := make(chan string)
	c :=  boring("boring!")  
        functions returning a channel
    
        joe := boring("Joe ")
        ann := boring("Ann ")
        */
	
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // receive expression is just a value
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
