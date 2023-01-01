package main

import "fmt"

func main() {
	// Create a set of strings
	set := make(map[string]struct{})

	// Add some items to the set
	set["1"] = struct{}{}
	set["2"] = struct{}{}
	set["3"] = struct{}{}

	// Check if an item is in the set
	_, ok := set["1"]
	fmt.Println("Is 1 in set? :", ok)

	_, ok = set["2"]
	fmt.Println("Is 2 in set? :", ok)

	// Remove an item from the set
	delete(set, "3")
	fmt.Println("Deleting 3 from set... ")
	// Iterate over the items in the set
	for key := range set {
		fmt.Println(key)
	}
}
