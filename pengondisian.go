package main

import "fmt"

func main() {
	// Menggunakan float64 untuk angka desimal
	price := 50000.0 // float64

	// Mengalikan harga dengan diskon
	if price > 300000 {
		price = price * 0.9 // 0.9 adalah float64
		fmt.Println(price)
		return
	}
	if price > 100000 {
		price = price * 0.95
		fmt.Println(price)
		return
	}
	fmt.Println(`Anda Tidak mendapatkan Discount, total biaya:`, price)
}
