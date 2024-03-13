// Given a positive integer n, find the pivot integer x such that:

// The sum of all elements between 1 and x inclusively equals the sum of all elements between x and n inclusively.
// Return the pivot integer x. If no such integer exists, return -1. It is guaranteed that there will be at most one pivot index for the given input.

// Example 1:
// Input: n = 8
// Output: 6
// Explanation: 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21.

// Example 2:
// Input: n = 1
// Output: 1
// Explanation: 1 is the pivot integer since: 1 = 1.

// Example 3:
// Input: n = 4
// Output: -1
// Explanation: It can be proved that no such integer exist.

package main

import "fmt"

func sumlist(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func pivotInteger(n int) int {
	numList := []int{}
	for i := 1; i <= n; i++ {
		numList = append(numList, i)
	}

	for i := 1; i <= n; i++ {
		if sumlist(numList[:i]) == sumlist(numList[i-1:]) {
			return numList[i-1]
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
}
