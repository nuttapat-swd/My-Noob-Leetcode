package utils

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
	if index >= len(nums) {
		return nil
	}
	tree := TreeNode{Val: nums[index]}
	tree.Left = PerformCreate(nums, 2*index+1)
	tree.Right = PerformCreate(nums, 2*index+2)
	return &tree
}
