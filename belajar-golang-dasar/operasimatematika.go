package main

import "fmt"

func main() {
	//operasi matematika
	var a = 28
	var b = 18
	var c = 1.2
	var d = 1.8
	var e = a + b
	var f = c + d

	fmt.Println("Hasil penjumlahan a + b adalah :", e)
	fmt.Println("Hasil penjumlahan c + d adalah :", f)
	// tidak bisa mengoperasikan antara tipe data yang berbeda, seperti a + c, karena a adalah int dan c adalah float64
	// fmt.Println("Hasil pengurangan a - b adalah :", a-c)

	//Augmented Assignment
	a += 2 // a = a + 2
	fmt.Println("Hasil augmented assignment a += 2 adalah :", a)
	a -= 2 // a = a - 2
	fmt.Println("Hasil augmented assignment a -= 2 adalah :", a)
	a *= 2 // a = a * 2
	fmt.Println("Hasil augmented assignment a *= 2 adalah :", a)
	a /= 2 // a = a / 2
	fmt.Println("Hasil augmented assignment a /= 2 adalah :", a)

	var i = 2025
	fmt.Println("Nilai i sebelum increment adalah :", i)
	// Increment or Unary Operator
	i++ // i = i + 1
	fmt.Println("Hasil increment i++ adalah :", i)
	i += 7
	fmt.Println("Hasil increment i += 7 adalah :", i)

}
