// Objective:Create a command-line interface (CLI) application in Go that allows users to input a block of text and then calculates the frequency of each word in the text. This project will help you understand and implement basic Go concepts like variables, loops, conditionals, maps, functions, and strings.
// Requirements:
// Input Text: Allow users to input a block of text.
// Process Text: Count the frequency of each word in the text.
// Display Frequencies: Display the word frequencies in a readable format.
// Exit: Allow the user to exit the application.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountFrequency(str string) map[string]int {
	word := strings.Fields(str)

	wc := make(map[string]int)
	// var frequency []string;
	for _, value := range word {
		_, matched := wc[value]
		if matched {
			wc[value] += 1
		} else {
			wc[value] = 1
		}
	}
	return wc
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the text \n")
	value, _ := reader.ReadString('\n')
	count := CountFrequency(value)
	fmt.Print(count)
	fmt.Print("Do you want to exit(y/n)")
    in, _ := reader.ReadString('\n')
    if in == "y" {
        os.Exit(0)
    }
}
