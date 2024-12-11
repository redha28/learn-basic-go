package main

import "fmt"

func main() {
	numbers := [...]int{5, 10, 15, 20, 25, 30}
	slice1 := numbers[3:6]
	slice2 := numbers[:3]
	fmt.Println(slice1, slice2)
	fmt.Println(len(numbers))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	daySlice := days[5:]
	daySlice[0] = "sabtu baru"
	daySlice[1] = "minggu baru"
	fmt.Println(daySlice)
	fmt.Println(days)

	daySlice2 := append(daySlice, "hari libur")
	daySlice2[0] = "on site"
	fmt.Println(daySlice2)
	fmt.Println(daySlice)
	fmt.Println(days)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
}