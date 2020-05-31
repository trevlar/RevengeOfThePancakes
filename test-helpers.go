package main

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
		OldSolveStack(testCase.ID, testCase.Input)
	}
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
