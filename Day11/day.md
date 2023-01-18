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

<p align="center">
 <img src="linkedlist.JPG?raw=true" alt="linear Data Structures" width="50%" height="50%" />
</p>
