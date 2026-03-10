package main

import "fmt"

func main() {
	var ipkakhir float32 = 3.9
	var skripsi = 3.6

	fmt.Println("Nilai IPK Akhir :", ipkakhir)
	var lulusipkakhir bool = ipkakhir >= 3.0
	var lulusskripsi bool = skripsi >= 3.0
	fmt.Println("Lulus IPK Akhir :", lulusipkakhir)
	fmt.Println("Lulus Skripsi :", lulusskripsi)

	if lulusipkakhir && lulusskripsi {
		fmt.Println("Selamat Anda Lulus")
	} else {
		fmt.Println("Maaf Anda Tidak Lulus")
	}
}
