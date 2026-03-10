package main

import "fmt"

func main() {
	// Array adalah tipe data yang dapat
	// menyimpan banyak nilai dengan tipe data yang sama,
	// dan memiliki ukuran tetap (fixed size).
	// Array memiliki indeks yang dimulai dari 0,
	// sehingga elemen pertama berada pada indeks 0,
	// elemen kedua pada indeks 1, dan seterusnya.

	// Deklarasi array dengan tipe data int dan ukuran 5
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Isi array numbers:", numbers[0], "+", numbers[1], "adalah", numbers[0]+numbers[1])
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Nilai pada indeks", i, "adalah", numbers[i])
	}

	var ips = [7]float32{
		3.5, 3.6, 3.7, 3.8, 3.9, 4.0,
	}

	// isi array tidak boleh lebih dari apa yang sudah ditentukan, namun jika kurang?
	//jika kurang dari total yang ditentukan maka tidak apa apa
	fmt.Println("Isi array ips:", ips)

	//FUNCTION ARRAY
	var lengthofips = len(ips)
	fmt.Println("Panjang array ips adalah :", lengthofips)
	fmt.Println("Nilai pada indeks 0 adalah:", ips[0])
	ips[0] = 7.8
	fmt.Println("Nilai pada indeks 0 setelah diubah adalah:", ips[0])

	//jika tidak mau menentukan jumlah secara pasti maka gunakan [...]
	// dan dia akan menyesuaikan dengan jumlah data yang kita masukkan di awal deklarasi array

	var ips2 = [...]float32{
		3.5, 3.6, 3.7, 3.8, 3.9, 4.0,
	}
	fmt.Println("Isi array ips2:", ips2)
	fmt.Println("Panjang array ips2 adalah :", len(ips2))
	//namun [...] hanya bisa digunakan saat deklarasi langsung array,
	//tidak bisa digunakan saat deklarasi array kosong dan kemudian diisi nilainya

	// DI GOLANG TIDAK ADA ISTILAH MENGHAPUS ARRAY, PALING ISI DATA KOSONG AJA
	//BEDA DENGAN JAVASCRIPT YANG BISA MENGHAPUS ARRAY DENGAN METHOD SPLICE ATAU POP
	// 3 ya tiga, gitu

}
