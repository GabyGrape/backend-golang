package main

import "fmt"

// type declarations adalah cara untuk membuat tipe data baru dengan menggunakan tipe data yang sudah ada
//

func main() {

	// membuat tipe data baru dengan menggunakan tipe data string bernama NoKTP
	//Tapi pertanyaannya kenapa harus bikin tipe data baru? kapan? bukankah di DB itu udah ke cover?
	type NoKTP string

	var ktpEko NoKTP = "111111"
	fmt.Println("No KTP Eko adalah :", ktpEko)

	var contoh string = "Ini adalah contoh string"
	//cara ubah tipe data nya ke typedate
	var contoh2 NoKTP = NoKTP(contoh)
	fmt.Println("Contoh adalah :", contoh)
	fmt.Println("Contoh 2 adalah :", contoh2)

}
