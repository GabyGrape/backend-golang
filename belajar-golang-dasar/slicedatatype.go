package main

func main() {
	// Slice adalah tipe data yang dapat menyimpan banyak nilai dengan tipe data yang sama,
	// namun memiliki ukuran yang dinamis (dynamic size).
	// Slice memiliki indeks yang dimulai dari 0, sehingga elemen pertama berada pada indeks 0,
	// elemen kedua pada indeks 1, dan seterusnya.

	//slice yang paling sering digunakan dalam golang
	//slice adalah potongan dari data array, dan size nyabisa nambah secara dinamis
	//slice selalu berhubungna dengan array, data yang akses sebagian atau seluruh data di array
	//setiap pembuatan slice itu akan membuat array baru, dan slice akan menunjuk ke array tersebut
	//slice itu seperti jendela yang bisa melihat sebagian atau seluruh data di array, dan bisa mengubah data di array tersebut
	//slice itu pintar dan manager pada array, walau dibalik dia bikin array baru tiap penambahan size

	//pointer, length, capacity
	//pointer adalah alamat memori dari array yang diakses oleh slice
	//length adalah jumlah elemen yang ada di slice
	//capacity adalah jumlah elemen yang bisa ditampung oleh slice, dan biasanya lebih besar atau sama dengan length

}
