package main

import (
	"fmt"
	"regexp"
	"strings"
)

// FindAndFlip determines if the stack of pancakes is ready to serve.
// This function calls FindFromTop to determine what pancake to flip next.
// When we find a pancake from the top of the stack, the group is flipped using FlipPancake.
// It counts each flip of the stack.
// If the stack is all happy-side up, then it's ready to serve.
// It returns the resulting stack and number of flips.
// As a precaution, if the stack gets flipped 100 times or more, then return out of the function.
func OldFindAndFlip(stack string) (string, int) {
	resultStack := stack
	numberOfFlips := 0
	limit := 100
	if isVerbose {
		fmt.Printf("stack: %s\n", stack)
	}
	for {
		pancakeToFlip := OldFindFromTop(resultStack)
		if pancakeToFlip == -1 || numberOfFlips >= limit {
			return resultStack, numberOfFlips
		}
		numberOfFlips = numberOfFlips + 1
		resultStack = OldFlipPancake(pancakeToFlip, resultStack)
		if isVerbose {
			if pancakeToFlip == 0 {
				fmt.Println("The top pancake was flipped over.")
			} else if pancakeToFlip+1 == len(resultStack) {
				fmt.Println("The entire pancake stack was flipped over.")
			} else {
				fmt.Printf("The top %d pancakes were flipped over.\n", pancakeToFlip+1)
			}
			fmt.Printf("Result: %s\n", resultStack)
		}
	}
}

// FindFromTop discovers the next pancake in the stack to flip.
// It iterates over the stack starting at the top pancake through til the bottom.
// When it finds a pancake that has a different side up, it returns the index of the previous pancake.
// If every pancake in the stack is happy-side up, then we return -1.
func OldFindFromTop(stack string) int {
	result := -1
	lastSide := ""
	allHappy := 0
	for _, c := range stack {
		side := string(c)
		if lastSide != "" && lastSide != side {
			return result
		}
		result = result + 1
		if side == happySide {
			allHappy = allHappy + 1
		}
		lastSide = side
	}
	if allHappy == len(stack) {
		return -1
	}
	return result
}

// FlipPancake flips over any pancakes before or up to the number of pancakes we want to flip.
// Flipping a pancake means that the character representing the pancake changes from + to - or from - to +.
// Any pancake that comes after the flipped group of pancakes remains unchanged.
func OldFlipPancake(numPancakes int, stack string) string {
	// a new string is returned with a numPancakes switched from the top.
	result := ""
	for i, c := range stack {
		side := string(c) // convert rune to string
		// ignore any character that is not the happy side or blank side.
		// validation occurs before this point.
		if side == happySide || side == blankSide {
			if i <= numPancakes {
				if side == blankSide {
					result = result + happySide
				} else {
					result = result + blankSide
				}
			} else {
				result = result + side
			}
		}
	}
	return result
}

// VerifyPancakeStack verifies that the composition and size of the stack of pancakes is valid and within limits
// The stack must be between 1 to 100 pancakes.
// Every character in the stack must be either a + or - symbol­.
func OldVerifyPancakeStack(stack string) (bool, error) {
	// If the stack has any characters other than a + or - it is considered invalid.
	r, _ := regexp.Compile("[^-+]")
	if r.MatchString(stack) == true {
		fmt.Printf("Invalid characters: %s. The string must consist of the characters '+' or '-'.\n", stack)
		return false, fmt.Errorf("invalid characters")
	}
	if len(stack) < 1 || len(stack) > 100 {
		fmt.Printf("Invalid stack size: %d. Must be between 1 to 100 pancakes tall.\n", len(stack))
		return false, fmt.Errorf("invalid stack size")
	}
	return true, nil
}

// SolveStack takes a stack of pancakes as a string and verify that the stack meets the criteria for this challenge.
// It will find and flip the stack and output the results.
func OldSolveStack(testID int, stack string) {
	isValid, err := OldVerifyPancakeStack(stack)
	if err != nil {
		fmt.Errorf("Invalid pancake stack: %v", err)
	} else if isValid {
		if isVerbose {
			fmt.Printf("\nSolving Case #%d\n", testID)
		}
		resultStack, numberOfFlips := OldFindAndFlip(stack)
		if isVerbose {
			if numberOfFlips == 0 {
				fmt.Printf("This stack was already happy-side up.\n")
			} else if numberOfFlips == 1 {
				fmt.Printf("Only %d flip to get this stack of pancakes happy-side up.\n", numberOfFlips)
			} else {
				fmt.Printf("It took %d flips to get this stack of pancakes happy-side up.\n", numberOfFlips)
			}
			fmt.Printf("Order up! %s\n", resultStack)
		}
		fmt.Printf("Case #%d: %d,\n", testID, numberOfFlips)
	}
}

// ConvertNewLine changes CRLF to LF
func OldConvertNewLine(text string) string {
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "\n", "", -1)
	return text
}
