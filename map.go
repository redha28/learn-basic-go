package main

import "fmt"

func main() {
	person := map[string]string {
		"name" : "eko",
		"address": "indonesia",
	}
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["title"] = "Golang"
	book["author"] = "eko"
	book["delete"] = "salah"

	fmt.Println(book)
	delete(book, "delete")
	
	fmt.Println(book)
}