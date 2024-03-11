package main

import (
	"fmt"
	"slices"
)

func intersection(nums1 []int, nums2 []int) []int {
	interectNums := []int{}
	for _, v := range nums1 {
		if slices.Contains(nums2, v) && !slices.Contains(interectNums, v) {
			interectNums = append(interectNums, v)
		}
	}
	return interectNums
}

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Println(intersection(num1, num2))

}
