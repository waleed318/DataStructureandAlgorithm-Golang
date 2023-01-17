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

// AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)

	fmt.Println("Node [1]: ", linkedList.headNode.property)
	fmt.Println("Node [2]: ", linkedList.headNode.nextNode.property)

}
