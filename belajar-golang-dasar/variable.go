package main

import "fmt"

func main() {

	//deklarasi variabel dengan tipe data string tanpa assign value
	var name string

	//assign value ke variabel name
	name = "Okta Gabriel Sinsaku Sinaga"
	fmt.Println("Nama saya adalah :", name)

	//name valuenya diubah
	name = "Gabriel Althea Buntaro"
	fmt.Println("Nama saya adalah :", name)

	//deklrasi variabel langsung assign value maka tipe data tidak perlu di state
	var name2 = "Elemonaremba"
	fmt.Println("Nama saya adalah :", name2)

	//kata kunci var di golang tidaklah wajib
	//:= dipakai untuk deklarasi variabel tanpa kata kunci var, tapi harus langsung assign value
	//dan tipe data akan otomatis terdeteksi oleh golang
	//dan jika ingin mengubah valuenya maka jangan gunakan := lagi, karena := hanya untuk deklarasi variabel baru,
	//  jika ingin mengubah value variabel yang sudah ada maka cukup dengan nama variabelnya saja
	avariabel3 := "Yang aku sayang adalah keluarga inti dan kekasihku"
	fmt.Println("Variabel 3 adalah :", avariabel3)

	//MULTIPLE VARIABLES
	var (
		firstName = "Okta"
		lastName  = "Gaby"
	//firstname declared but not used, may result to error, solution use _ for ignore unused variable
	)

	fmt.Println("Nama lengkap saya adalah :", firstName, lastName)
}
