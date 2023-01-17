package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	var node = Node{}
	node.property = 1
	linkedList.headNode = &node

	var node2 = Node{}
	node2.property = 2
	node2.nextNode = linkedList.headNode
	linkedList.headNode = &node2

	fmt.Println("Node [1]: ", linkedList.headNode.property)
	fmt.Println("Node [2]: ", linkedList.headNode.nextNode.property)

}
