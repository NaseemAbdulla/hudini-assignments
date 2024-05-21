// Objective:
// Learn how to send and receive values using channels.

// Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
package main

import "fmt"

func add(message chan string) {
	message <- "hello world"
}
func main() {
	messages := make(chan string)
	go add(messages)
	msg := <-messages
	fmt.Printf(msg)
}

// The main function should receive the message from the channel and print it.
// Hints:

// Use an unbuffered channel for simplicity.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use a buffered channel.

// Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

// -------------------------------------------------------------------------------------

// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
