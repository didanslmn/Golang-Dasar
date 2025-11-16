package main

import "fmt"

func main() {
	// Deklarasi dan Inisialisasi Array
	var buah [3]string
	buah[0] = "Mangga"
	buah[1] = "Apel"
	buah[2] = "Pisang"

	// Deklarasi langsung dengan inisialisasi
	warna := [4]string{"merah", "hijau", "biru", "kuning"}

	fmt.Println("Buah pertama:", buah[0])
	fmt.Println("Warna kedua:", warna[1])
	fmt.Println("Jumlah buah:", len(buah))
	fmt.Println("Jumlah warna:", len(warna))

	// iterasi array menggunakan for loop
	for i, b := range buah {
		fmt.Printf("Buah ke-%d: %s\n", i, b)
	}
}
