// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.
package main

import "fmt"

func add(message chan string) {
	message <- "hello"
}
func addd(message chan string) {
	message <- "world"
}

func main() {
	messages := make(chan string)
	go add(messages)
	go addd(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
