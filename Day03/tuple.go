package main

import "fmt"

func PowerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	// Create a tuple with two elements: a string and an integer
	Square, Cube := PowerSeries(2)

	// Print the tuple
	fmt.Println("Square: ", Square)
	fmt.Println("Cube: ", Cube)
}
