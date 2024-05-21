package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("hello", "world")
	format := "I just swapped %s with %s";
	fmt.Printf(format,a,b)
	// fmt.Printf("%s world", "hello")
}
