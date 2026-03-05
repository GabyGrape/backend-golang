package main

import "fmt"

func main() {
	// String adalah tipe data teks?
	fmt.Println("Ini adalah string")      // string literal
	fmt.Println(len("Ini adalah string")) // len untuk menghitung panjang string
	fmt.Println("Ini adalah string dengan tipe data string :", "Belajar Golang Dasar"[0])
}
