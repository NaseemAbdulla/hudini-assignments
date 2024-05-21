// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
package main

import (
	"fmt"
	"sync"
)

func route1(wt *sync.WaitGroup, ch chan string) {
	defer wt.Done()
	ch <- "first"
}
func route2(wt *sync.WaitGroup, ch chan string) {
	defer wt.Done()
	ch <- "second"
}
func route3(wt *sync.WaitGroup, ch chan string) {
	defer wt.Done()
	ch <- "third"
}
func main() {
	wt := new(sync.WaitGroup)
	channel := make(chan string)

	wt.Add(3)
	go route1(wt, channel)
	go route2(wt, channel)
	go route3(wt, channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	wt.Wait()

}
