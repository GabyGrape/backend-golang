// a file to test, that Go is error if have more than 1 main function in the same package
package main


//modul wajib untuk println atau menampilkan teks di GO.
//Eksplore lebih lanjut tentang module fmt di https://golang.org/pkg/fmt/
import "fmt"

func main() {
	fmt.Println("Hello Okta, greet the World?!")
}




// PS C:\Users\oktag\OneDrive\Desktop\Jualan\backend-golang\belajar-golang-dasar> go build
// # belajar-golang-dasar
// .\sample.go:8:6: main redeclared in this block
//         .\helloworld.go:7:6: other declaration of main
// PS C:\Users\oktag\OneDrive\Desktop\Jualan\backend-golang\belajar-golang-dasar> 

// Solution develop or learning, run one by one, use go run sample.go