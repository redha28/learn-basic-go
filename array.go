package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println(numbers[0], numbers[1], numbers[2])

	var names = [...]string {
		"Redha",
		"Definto",
	}
	fmt.Println(names)
}