package utils

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
