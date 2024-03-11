package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBinaryTree(nums []int) *TreeNode {
	return PerformCreate(nums, 0)
}

func PerformCreate(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == math.MinInt {
		return nil
	}
	tree := TreeNode{Val: nums[index]}
	tree.Left = PerformCreate(nums, 2*index+1)
	tree.Right = PerformCreate(nums, 2*index+2)
	return &tree
}

func scanValue(result *int, root *TreeNode, low int, high int) {
	if root == nil {
		return
	}

	if root.Val >= low && root.Val <= high {
		*result += root.Val
	}

	scanValue(result, root.Left, low, high)
	scanValue(result, root.Right, low, high)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result int
	scanValue(&result, root, low, high)
	return result
}
func main() {
	root := CreateBinaryTree([]int{10, 5, 15, 3, 7, math.MinInt, 18})
	fmt.Println(rangeSumBST(root, 7, 15))
}
