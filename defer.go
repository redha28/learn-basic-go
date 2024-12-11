package main

import "fmt"

// defer adalah memanggil sebuah function ketika eksekusi sebuah function selesai

func logging() {
	fmt.Println("Selesai Memanggil Application")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}