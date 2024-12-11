package main

import "fmt"

// function
// func sayHello(firstName string, lastName string) {
// 	fmt.Println("Hallo", firstName, lastName)
// }

// return function
// func namaFunction(parameter type) typeReturn {}
// func sayHello(name string) string {
// 	return "Hello " + name
// }

// multyple return function
// func tidak hanya mengembalikan 1 return tetapi juga bisa lebih dari 1
// func namaFunction(parameter type, parameter type) (string, string) {}
// func sayHello(firstName string, lastName string) (string, string) {
// 	return firstName, lastName
//  }

// named return values
// func sayHello() (firstName, middleName, lastName string) {
// 	firstName = "redha"
// 	middleName = "definto"
// 	lastName = "aja"
// 	return firstName, middleName, lastName
// }

// varargs
// sebuah parameter yang berada di posisi terakhir, memiliki kemampuan untuk dijadikan sebagai varargs
// varargs datanya bisa menerima lebih dari satu input atau anggap seperti array bedanya varargs kita bisa langsung inputkan datanya
// func sumAll(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	// firstName, lastName := sayHello("redha", "definto")
	
	// firstName, _ := sayHello("Redha", "definto")
	// fmt.Println("Hello", firstName, lastName)
	// fmt.Println("Hello", firstName)


	// firstName, middleName, lastName := sayHello()
	// fmt.Println("Hello", firstName, middleName, lastName)

	// result := sumAll(10, 10, 10, 10, 10)
	// numbers := []int{10,10,10,10,10}
	// result := sumAll(numbers...)

	// fmt.Println(result)

	sayHelloWithFilter("redha", spamFilter)
}