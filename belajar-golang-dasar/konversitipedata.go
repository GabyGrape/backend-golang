package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	//kalau jangkauannya terlalu bersar dibanding tujuan, maka tidak akan sesuai dengan nilai aslinya, karena terjadi overflow
	// dan akan kembali ke belakang atau batas bawah jadi looping dari bawah ke atas untuk memenuhi overflow
	// seperti int16 yang nerima int32, maka nilainya akan -3268
	fmt.Println("Nilai int32 :", nilai32)
	fmt.Println("Nilai int64 :", nilai64)
	fmt.Println("Nilai int16 :", nilai16)

	var name = "Elemonareremba"
	var e uint8 = name[0]   // e adalah byte yang merepresentasikan karakter pertama dari string name
	var estring = string(e) // estri adalah string yang dibuat dari byte e

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(estring)
}
