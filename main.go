package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// pancakes with a happy-side are defined as a + character.
// pancakes with a blank-side are defined as a - character.
const (
	happySide = "+"
	blankSide = "-"
)

// The terminal allows verbose mode which will log out the transformation of the pancake stack as it happens.
var isVerbose = false

// FindAndFlip determines if the stack of pancakes is ready to serve.
// This function calls FindFromTop to determine what pancake to flip next.
// When we find a pancake from the top of the stack, the group is flipped using FlipPancake.
// It counts each flip of the stack.
// If the stack is all happy-side up, then it's ready to serve.
// It returns the resulting stack and number of flips.
// As a precaution, if the stack gets flipped 100 times or more, then return out of the function.
func FindAndFlip(stack string) (string, int) {
	resultStack := stack
	numberOfFlips := 0
	limit := 100
	if isVerbose {
		fmt.Printf("stack: %s\n", stack)
	}
	for {
		pancakeToFlip := FindFromTop(resultStack)
		if pancakeToFlip == -1 || numberOfFlips >= limit {
			return resultStack, numberOfFlips
		}
		numberOfFlips = numberOfFlips + 1
		resultStack = FlipPancake(pancakeToFlip, resultStack)
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

// FlipPancake flips over any pancakes before or up to the number of pancakes we want to flip.
// Flipping a pancake means that the character representing the pancake changes from + to - or from - to +.
// Any pancake that comes after the flipped group of pancakes remains unchanged.
func FlipPancake(numPancakes int, stack string) string {
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
func VerifyPancakeStack(stack string) (bool, error) {
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

// NewTestCases will create the object to house our test cases and populate the slice.
func NewTestCases() TestCases {
	var testCases TestCases
	testCases.GetTestCases()
	return testCases
}

// TestCases is the structure that houses the testcases for input.
type TestCases struct {
	cases []TestCase
}

// GetTestCases will append the test cases for the input of this program.
func (t *TestCases) GetTestCases() {
	t.cases = append(t.cases, TestCase{1, "-", 1})
	t.cases = append(t.cases, TestCase{2, "-+", 1})
	t.cases = append(t.cases, TestCase{3, "+-", 2})
	t.cases = append(t.cases, TestCase{4, "+++", 0})
	t.cases = append(t.cases, TestCase{5, "---+-", 3})
}

// TestCase is a structure used in testcases for defining the default test cases.
type TestCase struct {
	ID             int
	Input          string
	ExpectedOutput int
}

// HandleTestCases will run the test cases for this program.
// The following conventions are applied for input and output.
// Input
// The first line of the input gives the number of test cases, T. T test cases follow. Each consists of one line
// with a string S, each character of which is either + (which represents a pancake that is initially happy side
// up) or ­ (which represents a pancake that is initially blank side up). The string, when read left to right,
// represents the stack when viewed from top to bottom.
// Output
// For each test case, output one line containing Case #x: y, where x is the test case number (starting
// from 1) and y is the minimum number of times you will need to execute the maneuver to get all the
// pancakes happy side up.
func HandleTestCases() {
	test := NewTestCases()
	for _, testCase := range test.cases {
		SolveStack(testCase.ID, testCase.Input)
	}
}

// This terminal program accepts arguments for various functionality.
func main() {
	args := os.Args
	if len(args) > 1 {
		var handleInput bool
		var handleTestCases bool
		stackInput := ""
		fileInput := ""
		for i := range args {
			switch args[i] {
			case "help", "-h":
				if len(args) > i+1 && args[i+1] == "-c" {
					ShowChallenge()
				} else {
					ShowHelpManual(args[0])
				}
			case "-i", "--input":
				handleInput = true
			case "-t", "--test":
				handleTestCases = true
			case "-s":
				if len(args) > i+1 {
					stackInput = args[i+1]
				}
			case "-f":
				if len(args) > i+1 {
					fileInput = args[i+1]
				}
			case "-v":
				isVerbose = true
			}
			if strings.Contains(args[i], "--stack=") {
				command := strings.Split(args[i], "=")
				if len(command) > 1 {
					stackInput = command[1]
				}
			} else if strings.Contains(args[i], "--file=") {
				command := strings.Split(args[i], "=")
				if len(command) > 1 {
					fileInput = command[1]
				}
			}
		}
		if handleTestCases {
			HandleTestCases()
		} else if stackInput != "" {
			SolveStack(1, stackInput)
		} else if fileInput != "" {
			HandleFile(fileInput)
		} else if handleInput {
			HandleInput()
		} else {
			ShowHelpManual(args[0])
		}
	} else {
		ShowHelpManual(args[0])
	}
}

// SolveStack takes a stack of pancakes as a string and verify that the stack meets the criteria for this challenge.
// It will find and flip the stack and output the results.
func SolveStack(testID int, stack string) {
	isValid, err := VerifyPancakeStack(stack)
	if err != nil {
		fmt.Errorf("Invalid pancake stack: %v", err)
	} else if isValid {
		if isVerbose {
			fmt.Printf("\nSolving Case #%d\n", testID)
		}
		resultStack, numberOfFlips := FindAndFlip(stack)
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

// HandleFile will take a file where each line is expected to be a new stack of pancakes to solve.
func HandleFile(filePath string) error {
	fmt.Println("readFileWithReadString")

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	testID := 1
	var line string
	for {
		line, err = reader.ReadString('\n')

		line = ConvertNewLine(line)

		SolveStack(testID, line)
		testID++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return nil
}

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
			SolveStack(testID, text)
			testID++
			fmt.Print("\n")
		}
	}
}

// ConvertNewLine changes CRLF to LF
func ConvertNewLine(text string) string {
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "\n", "", -1)
	return text
}

// ShowHelpManual will list the possible commands that this terminal program receives.
func ShowHelpManual(program string) {
	if strings.Contains(program, ".exe") {
		program = "go run main.go"
	}
	fmt.Printf(`usage: %s [-v] [-h] [--help] [-t] [--test] [-i]
           [--input] [--file=<file>] [-f file] [--stack=<stack>] [-s <stack>]

This terminal program is one of many possible solutions to the Revenge of The Pancakes challenge.

Creative liberty was taken in this challenge to embellish the UI with features that provide a somewhat user friendly experience.

This program can receive a stack of pancakes. The stack expected must only consist of the characters + and -.
The stack of pancakes must be between 1 to 100 pancakes tall. The + represents the happy-side covered in delicious chocolate chips. The - character represents the blank side without any apparent delight.

Selecting a group of pancakes is achieved by iterating over the stack of pancakes from top to bottom. That group is flipped over onto its opposite side. Each flipped group of pancakes leads to the discovery of another group until the entire stack is consolidated and entirely composed of pancakes that are happy-side up.

When the whole pancake stack is happy-side up, it is ready to serve.

The following test cases will run upon using the flags -t or --tests

+--------+--------------+
| Input  |    Output    |
+--------+--------------+
|  -     |  Case #1: 1  |
|  -+    |  Case #2: 1  |
|  +-    |  Case #3: 2  |
|  +++   |  Case #4: 0  |
|  --+-  |  Case #5: 3  |
+-------+---------------+

'mkpancakes help -c' will display the original challenge.`, program)
}

// ShowChallenge will the display the text from the original document for the challenge.
func ShowChallenge() {
	fmt.Println(`
Revenge of the Pancakes

This problem came from a Google Code Jam project that was made available online in 2016. Feel free to
use the language you feel most comfortable in, but Go is our preferred language if you are able to use it.

Problem

The Infinite House of Pancakes has just introduced a new kind of pancake! It has a happy face made of
chocolate chips on one side (the "happy side"), and nothing on the other side (the "blank side").

You are the head waiter on duty, and the kitchen has just given you a stack of pancakes to serve to a
customer. Like any good pancake server, you have X­ray pancake vision, and you can see whether each
pancake in the stack has the happy side up or the blank side up. You think the customer will be happiest if
every pancake is happy side up when you serve them.

You know the following maneuver: carefully lift up some number of pancakes (possibly all of them) from
the top of the stack, flip that entire group over, and then put the group back down on top of any pancakes
that you did not lift up. When flipping a group of pancakes, you flip the entire group in one motion; you do
not individually flip each pancake. Formally: if we number the pancakes 1, 2, ..., N from top to bottom, you
choose the top i pancakes to flip. Then, after the flip, the stack is i, i­1, ..., 2, 1, i+1, i+2, ..., N. Pancakes 1,
2, ..., inow have the opposite side up, whereas pancakes i+1, i+2, ..., N have the same side up that they
had up before.

For example, let's denote the happy side as + and the blank side as ­. Suppose that the stack, starting
from the top, is ­­+­. One valid way to execute the maneuver would be to pick up the top three, flip the
entire group, and put them back down on the remaining fourth pancake (which would stay where it is and
remain unchanged). The new state of the stack would then be ­++­. The other valid ways would be to
pick up and flip the top one, the top two, or all four. It would not be valid to choose and flip the middle two
or the bottom one, for example; you can only take some number off the top.

You will not serve the customer until every pancake is happy side up, but you don't want the pancakes to
get cold, so you have to act fast! What is the smallest number of times you will need to execute the
maneuver to get all the pancakes happy side up, if you make optimal choices?

Input
The first line of the input gives the number of test cases, T. T test cases follow. Each consists of one line
with a string S, each character of which is either + (which represents a pancake that is initially happy side
up) or ­ (which represents a pancake that is initially blank side up). The string, when read left to right,
represents the stack when viewed from top to bottom.

Output
For each test case, output one line containing Case #x: y, where x is the test case number (starting
from 1) and y is the minimum number of times you will need to execute the maneuver to get all the
pancakes happy side up.

Limits
1 ≤ T ≤ 100.
Every character in S is either + or ­.

Small dataset
1 ≤ length of S ≤ 10.

Large dataset
1 ≤ length of S ≤ 100.

Sample

+--------+--------------+
| Input  |    Output    |
+--------+--------------+
|  -     |  Case #1: 1  |
|  -+    |  Case #2: 1  |
|  +-    |  Case #3: 2  |
|  +++   |  Case #4: 0  |
|  --+-  |  Case #5: 3  |
+-------+---------------+

In Case #1, you only need to execute the maneuver once, flipping the first (and only) pancake.

In Case #2, you only need to execute the maneuver once, flipping only the first pancake.

In Case #3, you must execute the maneuver twice. One optimal solution is to flip only the first pancake,
changing the stack to ­­, and then flip both pancakes, changing the stack to ++. Notice that you cannot
just flip the bottom pancake individually to get a one­move solution; every time you execute the
maneuver, you must select a stack starting from the top.

In Case #4, all of the pancakes are already happy side up, so there is no need to do anything.

In Case #5, one valid solution is to first flip the entire stack of pancakes to get +­++, then flip the top
pancake to get ­­++, then finally flip the top two pancakes to get ++++.`)
}
