package main

// func twoSum(nums []int, target int) []int {
// 	var output []int
// 	tmp := 0
// 	for i, v := range nums {
// 		tmp = v
// 		start_idx := 0
// 		output = append(output, i)
// 		fmt.Println("Starter", v)
// 		for start_idx < len(nums) {
// 			fmt.Println("new start", start_idx)
// 			fmt.Println("temp", tmp)
// 			for j, vl := range nums[start_idx:] {
// 				j += start_idx
// 				fmt.Print(tmp, vl, tmp+vl)
// 				fmt.Print(" - ")
// 				fmt.Print(i, j)
// 				fmt.Print(" - ")

// 				if (i == j) || (tmp+vl > target) {
// 					fmt.Println("Skip")
// 					continue
// 				} else if tmp+vl == target {
// 					fmt.Println("Target !")
// 					tmp += vl
// 					output = append(output, j)
// 					break
// 				} else {
// 					fmt.Println("include")
// 					tmp += vl
// 					break
// 					// output = append(output, j)
// 				}
// 			}
// 			if tmp == target {
// 				break
// 			}
// 			start_idx += 1
// 			tmp = v
// 		}

// 		if tmp == target {
// 			break
// 		} else {
// 			output = []int{}
// 		}
// 	}
// 	sort.Ints(output)
// 	fmt.Println(output)
// 	return output
// }

func twoSum(nums []int, target int) []int {
	var output []int
	// sort.Ints(nums)
	for i, v := range nums {
		tmp := v
		output = append(output, i)
		for j, vl := range nums {
			if i == j {
				continue
			} else if tmp+vl == target {
				tmp += vl
				output = append(output, j)
				break
			}
		}
		if tmp == target {
			break
		} else {
			output = []int{}
		}
	}
	// fmt.Println(output)
	return output
}

func main() {
	twoSum([]int{1, 3, 4, 2}, 6)
}
