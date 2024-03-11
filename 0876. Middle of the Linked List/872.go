package main

import "fmt"

// Definition for singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

func (n *Node) add(val int) {
	newNode := &Node{Val: val}
	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
	return
}

func middleNode(head *Node) *Node {
	if head == nil {
		return head
	}
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	head := &Node{Val: 1}
	head.add(2)
	head.add(3)
	head.add(4)
	head.add(5)

	a := middleNode(head)
	fmt.Println("/////////")
	fmt.Println(a)

}

// 1 2 3 4 5 6 7 8 9 0 11 12
//                           f
//             s
