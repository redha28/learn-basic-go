package main

import "fmt"

func main() {

	// model pertama
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke", counter)
	// 	counter++
	// }
	
	// for counter := 1; counter <= 10; counter++ {
			// fmt.Println("Perulangan Ke", counter)
	// }
	// fmt.Println("selesai")
	// names := []string{"eko", "kurniawan", "khanedy"}
	// for idx, name := range names {
	// 	fmt.Println("idx", idx, "=", name)
	// }
	// for _, name := range names {
	// 	fmt.Println(name)
	// }


	// break
	// ketika ketemu kata kunci break pada perulangan maka perulangan akan di berhentikan
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		fmt.Println("Perulangan di berhentikan")
	// 		break
	// 	}
	// 	fmt.Println("Perulangan ke", i)
	// }
	
	// continue
	// ketika ketemu kata kunci continue pada perulangan maka perulangan akan menghentikan perulangan saat ini dan melanjutkan perulangan berikutnya
	
	for i := 0; i < 10; i++ {
		if i %2 == 0 {
			fmt.Println("angka ini genap")
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}