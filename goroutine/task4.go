// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.
package main

func sum(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}
func main() {
	c := make(chan int)
	go sum(c)
	for i := range c {
		println(i)

	}

}
