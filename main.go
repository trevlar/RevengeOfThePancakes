package main

import (
	"fmt"
)

const (
	happySide = "+"
	blankSide = "-"
)

func FindAndFlip(stack string) (string, int) {
	resultStack := stack
	numberOfFlips := 0
	limit := 100
	for {
		pancakeToFlip := FindFromTop(resultStack)
		if pancakeToFlip == -1 || numberOfFlips >= limit {
			return resultStack, numberOfFlips
		}
		numberOfFlips = numberOfFlips + 1
		resultStack = FlipPancake(pancakeToFlip, resultStack)
	}
}

func FindFromTop(stack string) int {
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

func FlipPancake(numPancakes int, stack string) string {
	result := ""
	for i, c := range stack {
		side := string(c)
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

func verifyPancakes(stack string) {
	// Limits
	// 1 ≤ T ≤ 100.
	// Every character in S is either + or -­.
	// Small dataset
	// 1 ≤ length of S ≤ 10.
	// Large dataset
	// 1 ≤ length of S ≤ 100.
}

func main() {
	fmt.Printf("Pancakes made happy-side up\n")
	// Case #3: 2
	results := FindFromTop("+-")
	if results != 0 {
		fmt.Errorf("Found wrong pancake from stack to flip, got %d, want: %d.", results, 0)
	}
	// Case #5: 3
	// resultStack, numberOfFlips := FindAndFlip("--+-")
	// if resultStack != "++++" && numberOfFlips != 1 {
	// 	fmt.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++++", 3)
	// }
}

// pancakes to serve
// x-ray pancake vision happy side or blank side up
// Every pancake as happy side up

// lift some number of pancakes from stack
// flip entire group
// put group back down on top of any pancakes unlifted
// P: 1, 2, ..., N Top to bottom
// choose top i pancakes to flip
// After flip the stack is i, i-1, ..., 2, 1, i+1, i+2, ..., N.
// Pancakes 1, 2, ..., i have opposite side up
// Pancakes i+1, i+2, ..., N have same side up that they had before
// happy side = +
// blank side = -
// example (unit tests)
// --+-
// 1. pick up top three, flip entire group, put them back down on remaining fourth pancake
// New state is -++-
// 2. Puck up and flip top one, top two, or all four.
// Not valid to choose and flip the middle two or the bottom. Can only take N off of top.

// Don't serve until every pancake is happy side up.
// You don't want pancakes to get cold, so act fast!
// What is the lowest number of times you need to execute a maneuver to get all the pancakes happy side up. Given optimal choices.

//  Input
// The first line of the input gives the number of test cases, T. T test cases follow. Each consists of one line
// with a string S, each character of which is either + (which represents a pancake that is initially happy side
// up) or ­ (which represents a pancake that is initially blank side up). The string, when read left to right,
// represents the stack when viewed from top to bottom.

// Output
// For each test case, output one line containing Case #x: y, where x is the test case number (starting
// from 1) and y is the minimum number of times you will need to execute the maneuver to get all the
// pancakes happy side up.

// Limits
// 1 ≤ T ≤ 100.
// Every character in S is either + or -­.
// Small dataset
// 1 ≤ length of S ≤ 10.
// Large dataset
// 1 ≤ length of S ≤ 100.

// Input
// 5
// ­-
// ­-+
// +­-
// +++
// ­­--+-­

// Output
// Case #1: 1
// Case #2: 1
// Case #3: 2
// Case #4: 0
// Case #5: 3

// In Case #1, you only need to execute the maneuver once, flipping the first (and only) pancake.
// In Case #2, you only need to execute the maneuver once, flipping only the first pancake.
// In Case #3, you must execute the maneuver twice. One optimal solution is to flip only the first pancake,
// changing the stack to ­­, and then flip both pancakes, changing the stack to ++. Notice that you cannot
// just flip the bottom pancake individually to get a one­move solution; every time you execute the
// maneuver, you must select a stack starting from the top.
// In Case #4, all of the pancakes are already happy side up, so there is no need to do anything.
// In Case #5, one valid solution is to first flip the entire stack of pancakes to get +­++, then flip the top
// pancake to get ­­++, then finally flip the top two pancakes to get ++++.
