package main

import "fmt"

// When calling a function with params with the same type, we can declare as
/// func add (x, y int) int {}

func add(x int, y int) int {
	return x + y
}

func main() {
	// Declare a param in go by 2 ways
	/// 1st
	//// var result int = add (22, 23)

	/// 2nd
	result := add(22, 23)
	fmt.Println(result)

}
