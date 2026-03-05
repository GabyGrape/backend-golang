package main

import "fmt"

//constanta adalah variabel yang nilainya tidak bisa diubah setelah dideklarasikan

func main() {
	const mySurename = "Sinaga"
	const myFavouriteThings = "Makanan, Minuman, dan Keluarga"

	fmt.Println("Nama belakang saya adalah :", mySurename)
	fmt.Println("Hal yang saya sukai adalah :", myFavouriteThings)

	//will be error if try to assign a new value to a constant
	//mySurename = "Sinsaku" // error: cannot assign to mySurename (untyped string constant "Sinaga")

	//assign or declare multiple constant variabel

	const (
		food   = "Bakso"
		drinks = "Es Teh"
	)
	fmt.Println("Makanan favorit saya adalah :", food)
	fmt.Println("Minuman favorit saya adalah :", drinks)
}
