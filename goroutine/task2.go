// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.
package main

import (
	"fmt"
	// "sync"
	"time"
)

type WaitGroup struct {
	count int
}

func (wg *WaitGroup) add() {
	wg.count++
}
func (wg *WaitGroup) done() {
	wg.count--
}
func (wg *WaitGroup) wait() {
	for {
		if wg.count == 0 {
			break
		}
	}
}

func f(wg *WaitGroup) {
	for i := 1; i <= 5; i++ {

		fmt.Println(i, "route", wg)
		time.Sleep(time.Second)

	}
	defer wg.done()
}
func main() {
	wg := WaitGroup{}

	// wg := new(sync.WaitGroup)

	// go f("go route 1", wg)
	// go f("go route 2", wg)
	// go f("go route 3", wg)
	for i := 0; i < 3; i++ {
		wg.add()
		go f(&wg)

	}
	wg.wait()
	// time.Sleep(time.Second)
	// fmt.Printf("hello")
	//fmt.Println("from,")

}
