# Revenge Of The Pancakes

This repository houses a solution to a problem that came from a Google Code Jam project that was made available online in 2016. I chose to use the programming language of Go to solve this problem.

I overengineered my original solution. I've simplified the code to do exactly as the problem asks.

The solution encompasses iterating over the string of happy and blank side pancakes. There are two conditions in the loop. The first condition verifies if this is the last pancake in the stack, and it is blank. The second looks to see if the next pancake in the stack is happy-side up, and the prior is blank-side up. If it finds either case fulfilled, it increases the flips by one.

The final check happens after the loop has completed. If there is more than one flip and the first is happy-side up or if it required two or more flips, then return the number of flips plus one. Otherwise, return the flips.

## Running the Tests

To see the results of the tests run the following.

`go test`

### Prerequisites

This program was built and compiled using go 1.13.4.

## Authors

* **Trevor Carlston** - *Initial work* - [Trevlar](https://github.com/trevlar)

