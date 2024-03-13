package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) add(val int) {
	newNode := &ListNode{Val: val}
	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
	return
}

func (n *ListNode) print() {
	fmt.Print(n.Val)
	if n.Next != nil {
		n.Next.print()
	}
	fmt.Println("")
	return
}

// Fail T T
// func (n *ListNode) listNegative(result *[]int) {
// 	if n.Val < 0 {
// 		*result = append(*result, n.Val)
// 		if n.Next != nil {
// 			*n = *n.Next
// 		} else {
// 			*n = ListNode{}
// 			return
// 		}
// 	}
// 	if n.Next != nil {
// 		n.Next.listNegative(result)
// 	}
// 	return
// }

// func removeZeroSumSublists(head *ListNode) *ListNode {
// 	negativeValue := []int{}
// 	head.listNegative(&negativeValue)
// 	// fmt.Println(negativeValue)
// 	// head.print()
// 	// head.print()
// 	return head
// }

func removeZeroSumSublists(head *ListNode) *ListNode {
	sumMap := make(map[int]*ListNode)
	tmp := &ListNode{Val: 0, Next: head}

	sum := 0
	sumMap[sum] = tmp
	curr := tmp.Next
	for curr != nil {
		sum += curr.Val
		sumMap[sum] = curr
		curr = curr.Next
	}

	sum = 0
	curr = tmp
	for curr != nil {
		sum += curr.Val
		curr.Next = sumMap[sum].Next
		curr = curr.Next
	}

	return tmp.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.add(2)
	head.add(3)
	head.add(-3)
	head.add(4)
	removeZeroSumSublists(head)
}
