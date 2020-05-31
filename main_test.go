package main

import (
	"testing"
)

func TestFindAndFlipPancake(t *testing.T) {
	stack := "-"
	actual := findAndFlip(stack)
	expected := 1
	if actual != expected {
		t.Errorf("Test case 1 failed for stack `%v`, got %v, want: %v.", stack, actual, expected)
	}

	stack = "-+"
	actual = findAndFlip(stack)
	expected = 1
	if actual != expected {
		t.Errorf("Test case 2 failed for stack `%v`, got %v, want: %v.", stack, actual, expected)
	}

	stack = "+-"
	actual = findAndFlip(stack)
	expected = 2
	if actual != expected {
		t.Errorf("Test case 3 failed for stack `%v`, got %v, want: %v.", stack, actual, expected)
	}

	stack = "+++"
	actual = findAndFlip(stack)
	expected = 0
	if actual != expected {
		t.Errorf("Test case 4 failed for stack `%v`, got %v, want: %v.", stack, actual, expected)
	}

	stack = "---+-"
	actual = findAndFlip(stack)
	expected = 3
	if actual != expected {
		t.Errorf("Test case 5 failed for stack `%v`, got %v, want: %v.", stack, actual, expected)
	}
}
