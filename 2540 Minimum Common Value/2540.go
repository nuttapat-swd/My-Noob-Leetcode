package main

import (
	"fmt"
)

// func checkCompil(value int, list []int) bool {
// 	output := false

// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return output
// }

// func getCommon(nums1 []int, nums2 []int) int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)
// 	less1 := nums1[0]
// 	less2 := nums2[0]
// 	isCompil1 := checkCompil(less1, nums2)
// 	isCompil2 := checkCompil(less2, nums1)

//		if isCompil1 && !isCompil2 {
//			return less1
//		} else if !isCompil1 && isCompil2 {
//			return less2
//		} else if isCompil1 && isCompil2 {
//			less := less1
//			if less > less2 {
//				less = less2
//			}
//			return less
//		} else {
//			return -1
//		}
//	}

// // Map
// func getCommon(nums1 []int, nums2 []int) int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)
// 	interectNums := []int{}
// 	for _, v := range nums1 {
// 		if slices.Contains(nums2, v) {
// 			interectNums = append(interectNums, v)
// 		}
// 	}
// 	if len(interectNums) > 0 {
// 		return interectNums[0]
// 	}
// 	return -1
// }

func getCommon(nums1 []int, nums2 []int) int {
	index1, index2 := 0, 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			return nums2[index2]
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			index1++
		}
	}
	return -1
}

func main() {
	result := getCommon([]int{6, 18, 18, 28, 34, 37, 39, 50, 52, 54, 62, 63, 65, 66, 75, 80, 97, 98}, []int{10, 19, 27, 33, 40, 41, 43, 56, 61, 69, 72, 78, 79, 82, 88, 91, 94})
	fmt.Println(result)
}
