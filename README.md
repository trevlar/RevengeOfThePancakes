# RevengeOfThePancakes

This repository houses a solution to a problem that came from a Google Code Jam project that was made available online in 2016. I chose to use the programming language of Go to solve this problem.

There are a variety of other ways that this problem could be solved. This simple solution worked to solve the code challenge while providing a variety of methods of input from the terminal.

A possible solution I've considered is using numbers to represent groups of pancakes. Using numbers for groups of pancakes is especially useful for smaller data sets. In such a scenario, the stack for --+- would be processed using -100. Using this number would make it easier to construct an algorithm that would quickly determine an optimal solution. This system can be useful for stacks up to 16 characters long. When we solve a stack the size of 10 characters, we can switch to using hexadecimal base 16 numbering system. In this scenario the stack for --+-+++++++++++- would be processed using 100b0 Once it got beyond 16 it would become difficult rely on this method. We would need another method to determine an optimal solution.

Optimization can be accomplished using go channels to process multiple possible solutions for multiple stacks at the same time. I may see a speed increase by consolidating portions of the code where iteration over the stack takes place. I could add profiling and identify the slowest parts of the program and rewrite them for speed. 

However, I might not be going for performance. I may be writing for simplicity to make it easier for other developers to pick up where I left off. In such a case, I would use more types and break up parts of my program into separate files and modules by their responsibility. Separating the code into modules would reduce the cognitive load when learning the codebase. I would increase code reuse and simplify where possible. I would increase the number of comments to clarify confusing portions of my code.

I could add more unit tests to the test framework to cover all of the methods I wrote. The tests verify that the code work as designed and helps me identify any edge cases or assumptions that my code is making about input and output.

I know I can add better error handling to the code and provide better logging for these errors.

## Getting Started

In the project file run the following command in a terminal window to run the program.

`go run main.go -t`

There are various flags to test different portions of the program.

`-t` Run the test cases from the original challenge.
`-f <filename>` Read a file with each line expected to be a stack of pancakes.
`-s <stack>` Read an argument as a stack of pancakes to solve.
`-i` Persistent terminal input where different stacks of pancake can be tested.
`-h` Show help


```
go run main.go -t
go run main.go pancake-test.txt
go run main.go -s -+-+-++++-+------+
go run main.go -i
```

Enjoy!

### Prerequisites

This program was built and compiled using go 1.13.4.

### Installing

I'm not sure why you'd need to install this program. It's a solution to a challenge.

`go install`

After installing you'll be able to use the terminal to run the program to make pancake stacks.

`RevengeOfThePancakes -t`

## Running the tests

`go test`

## Authors

* **Trevor Carlston** - *Initial work* - [Trevlar](https://github.com/trevlar)

