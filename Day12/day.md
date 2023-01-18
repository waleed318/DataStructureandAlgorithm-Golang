# Linked list
A linked list is a data structure that consists of a sequence of elements, each containing a reference (or "link") to the next element in the list. The elements are commonly called nodes and the reference is usually implemented as a pointer. The first element in the list is called the head and the last element is called the tail. The last element's reference is typically set to nil to indicate the end of the list.

Linked lists are commonly used in computer science because they allow for efficient insertion and deletion of elements in the middle of the list, as well as efficient traversal of the list without the need for random access. They do, however, have the disadvantage of requiring more memory than arrays because of the additional pointer stored with each element.

## The AddToHead method
The AddToHead method adds the node to the start of the linked list. The
AddToHead method of the LinkedList class has a parameter integer property. The
property is used to initialize the node.

## The IterateList method
The IterateList method of the LinkedList class goes through the list starting from the headNode and prints the property of each node it encounters.

## The LastNode method
The LastNode method of the LinkedList class returns the last node in the list by traversing through the list and checking if the nextNode property is nil.

## The insertToEnd method
The insertToEnd method of the LinkedList class adds a node to the end of the list by first finding the current last node and then setting the nextNode property of that node to the new node being added.

## The NodeWithValue method
In the following code snippet, the NodeWithValue method of LinkedList returns the
node with the property value. The list is traversed and checked to see whether the
property value is equal to the parameter property

## The AddBetween method
The AddBetween method adds the node after a specific node. The AddBetween method of
LinkedList has nodeProperty and property parameters. A node with
the nodeProperty value is retrieved using the NodeWithValue method.

## The DeleteNode method
This method takes a single parameter, the property of the node that should be deleted. It iterates through the linked list, keeping track of the current node and the previous node. When it finds the node with the matching property, it updates the nextNode pointer of the previous node to skip over the node that should be deleted. If the node to be deleted is the headNode, it updates headNode to point to the next node in the list.

<p align="center">
 <img src="linkedlist.JPG?raw=true" alt="linear Data Structures" width="50%" height="50%" />
</p>
