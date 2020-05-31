package main

import (
	"fmt"
	"testing"
)

func TestFindAndTop(t *testing.T) {

	// Case #1: 1
	resultStack, numberOfFlips := OldFindAndFlip("-")
	if resultStack != "-" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "+", 1)
	}
	// Case #2: 1
	resultStack, numberOfFlips = OldFindAndFlip("+-")
	if resultStack != "++" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++", 1)
	}
	// Case #3: 2
	resultStack, numberOfFlips = OldFindAndFlip("-+")
	if resultStack != "++" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++", 2)
	}
	// Case #4: 0
	resultStack, numberOfFlips = OldFindAndFlip("+++")
	if resultStack != "+++" && numberOfFlips != 0 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "+++", 0)
	}
	// Case #5: 3
	resultStack, numberOfFlips = OldFindAndFlip("--+-")
	if resultStack != "++++" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++++", 3)
	}
	resultStack, numberOfFlips = OldFindAndFlip("+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+")
	if resultStack != "+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++" {
		t.Errorf("Found wrong pancake from stack to flip, got %s, %d want: %s %d.", resultStack, numberOfFlips, "++++", 3)
	} else {
		fmt.Printf("Awesome %s took %d flips", resultStack, numberOfFlips)
	}
}

func TestOldVerifyPancakeStack(t *testing.T) {
	// Case #1: 1
	results, _ := OldVerifyPancakeStack("-")
	if results != true {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, true)
	}
	// Case #2: 1
	results, _ = OldVerifyPancakeStack("-+")
	if results != true {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, true)
	}
	// Case #3: 2
	results, _ = OldVerifyPancakeStack("+-")
	if results != true {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, true)
	}
	// Case #4: 0
	results, _ = OldVerifyPancakeStack("+++")
	if results != true {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, true)
	}
	// Case #5: 3
	results, _ = OldVerifyPancakeStack("--+-")
	if results != true {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, true)
	}
	results, _ = OldVerifyPancakeStack("--a+-")
	if results != false {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, false)
	}
	results, _ = OldVerifyPancakeStack("--=+-")
	if results != false {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, false)
	}
	results, _ = OldVerifyPancakeStack("-- +-5")
	if results != false {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, false)
	}
	results, _ = OldVerifyPancakeStack("+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+-+-+-+------++++++-+-")
	if results != false {
		t.Errorf("Found wrong pancake from stack to flip, got %t, want: %t.", results, false)
	}
}
