package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

		line = OldConvertNewLine(line)

		OldSolveStack(testID, line)
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
