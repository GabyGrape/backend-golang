package main

import "fmt"

func main() {
	//operasi matematika

	var j = 10
	j++ // j = j + 1
	fmt.Println("Hasil increment j++ adalah :", j)
	j++ // j = j + 1
	fmt.Println("Hasil increment j++ adalah :", j)
	j--
	fmt.Println("Hasil decrement j-- adalah :", j)

	//OPERASI PERBANDINGAN membandingkan dua buah nilai atau data, dan menghasilkan boolean (true or false)
	var name = "Gaby"
	var name2 = "Gaby"
	var name3 = "Gaby2"
	var result = name != name2
	var first = 10
	var second = 20
	var result2 = first > second
	fmt.Println("Hasil perbandingan first > second adalah :", result2) // false

	fmt.Println("Hasil perbandingan name == name2 adalah :", name == name2) // true
	fmt.Println("Hasil perbandingan name == name3 adalah :", name == name3) // false
	fmt.Println("Hasil perbandingan name != name2 adalah :", result)        // true
}
