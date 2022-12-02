package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var intList list.List
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
