package main

// This terminal program accepts arguments for various functionality.
func main() {
	handleCommandLine()
}

// pancakes with a happy-side are defined as a + character.
// pancakes with a blank-side are defined as a - character.
const (
	happySide = "+"
	blankSide = "-"
)

func findAndFlip(stack string) (flips int) {
	lastSide := ""
	for i, c := range stack {
		side := string(c)

		lastPancakeBlank := side == blankSide && i == len(stack)-1
		pancakeSwitched := side == happySide && lastSide == blankSide

		if lastPancakeBlank || pancakeSwitched {
			flips++
		}
		lastSide = side
	}
	if flips > 0 && string(stack[0]) == happySide || flips > 1 {
		return flips + 1
	}
	return flips
}
