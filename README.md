# Revenge Of The Pancakes

This repository houses a solution to a problem that came from a Google Code Jam project that was made available online in 2016. I chose to use the programming language of Go to solve this problem.

I overengineered my original solution. I've simplified the code to do exactly as the problem asks.

The solution encompasses iterating over the string of happy and blank side pancakes. There are two checks the first looks to see if this is the last pancake in the stack is it blank. The second looks to see if the next pancake in the stack is happySide and the prior is blankSide. If either of these cases are true it will increase the flips by one.

The final check happens after it has iterated over the complete string. If there are more than one flip and the first is a happy side or if there is more than one flips then returns the flips and add one. Otherwise return the flips.

## Running the Tests

To see the results of the tests run the following.

`go test`

### Prerequisites

This program was built and compiled using go 1.13.4.

## Authors

* **Trevor Carlston** - *Initial work* - [Trevlar](https://github.com/trevlar)

