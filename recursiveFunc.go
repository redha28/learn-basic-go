package main

import "fmt"

// func with loop
// func factorialLoop(value int) int {
// 	result := 1

// 	for i := value; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

// func with recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		// fmt.Print(value)
		return value * factorialRecursive(value - 1)
	}
}



func main() {
	fmt.Print(factorialRecursive(10))
}