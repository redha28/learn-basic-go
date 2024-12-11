package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
}

// panic adalah function yang bisa kita gunakan untuk menghentikan program
// recover adalah funtion yang bisa kita gunakan untuk menangkap data panic

func runApp(error bool) {
	endApp()
	if error {
		panic("Error")
	}
}

func main() {
	runApp(true)
}