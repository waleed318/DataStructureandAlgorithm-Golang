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

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) insertToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) AddBetween(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	count := 0
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println("Node [", count, "]: ", node.property)
		count += 1
	}
}
func (linkedList *LinkedList) DeleteNode(property int) {
	var currentNode, previousNode *Node
	for currentNode = linkedList.headNode; currentNode != nil; currentNode = currentNode.nextNode {
		if currentNode.property == property {
			fmt.Println("Deleting node", property, "...")
			if previousNode != nil {
				previousNode.nextNode = currentNode.nextNode
			} else {
				linkedList.headNode = currentNode.nextNode
			}
			break
		}
		previousNode = currentNode
	}
}
func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.insertToEnd(4)
	linkedList.insertToEnd(8)
	linkedList.insertToEnd(16)
	linkedList.AddBetween(1, 3)
	linkedList.IterateList()
	linkedList.DeleteNode(3)
	linkedList.IterateList()
}
