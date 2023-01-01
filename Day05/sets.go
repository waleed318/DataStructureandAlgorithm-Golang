// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

import "fmt"

// importing fmt package

//Set class
type Set struct {
	integerMap map[int]bool
}

//create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

//deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

// main method
func main() {
	var set *Set
	set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
	set.DeleteElement(1)
	fmt.Println(set)
}
