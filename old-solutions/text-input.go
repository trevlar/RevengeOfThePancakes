package main

import (
	"bufio"
	"fmt"
	"os"
)

// HandleInput will run the persistent terminal program. It enables manual entry of stacks of pancakes.
func HandleInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter pancake stack")
	fmt.Println("Use the characters '+' and '-' to symbolize a pancake happy-side and blank-side up")
	fmt.Println("Press Enter to submit your order.")
	fmt.Println("Type 'quit' or exit to 'stop' the program.")
	fmt.Println("---------------------")
	var testID = 1
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = ConvertNewLine(text)
		if text != "" {
			if text == "exit" || text == "quit" ||
				text == "e" || text == "q" {
				break
			}
			flips := findAndFlip(text)

			fmt.Println(text, flips)

			testID++
			fmt.Print("\n")
		}
	}
}
