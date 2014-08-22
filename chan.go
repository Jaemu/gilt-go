package main

import (
	"fmt"
	"log"
	"time"
)

func Greet(name string, response_chan chan string) {
	time.Sleep(3 * time.Second)
	greeting := fmt.Sprintf("Greetings, %s!\n", name)
	response_chan <- greeting
	log.Print("Greeting function is sleeping")
	time.Sleep(3 * time.Second)
}

/* 'select' keyword
*	reads from the first channel that sends a value
* if multiple channels return at same time,
* chooses pseudorandomly
 */
func main() {
	cs := make(chan string)
	cs1 := make(chan string)
	go Greet("Alice", cs)
	go Greet("Bob", cs1)
	greeting := <-cs
	bobGreeting := <-cs1
	log.Print(bobGreeting)
	time.Sleep(10 * time.Second)
	log.Print(greeting)
}
