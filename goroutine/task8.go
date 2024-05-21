// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
package main

import (
	"fmt"
)

func b(ch chan string) {
	ch <- "first message"
}
func d(ch chan string) {
	ch <- "fsecond message"
}
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go b(channel1)
	go d(channel2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-channel1:
			fmt.Println("hi,", msg)
		case msg1 := <-channel2:
			fmt.Println("hi,", msg1)
		}
	}
}
