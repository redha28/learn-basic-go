package main

import "fmt"

type Product struct {
	Name, Size string
	stock      int
}

func main() {
	 var Ciki Product

	Ciki.Name = "Potato"

	fmt.Println(Ciki.Name)
}