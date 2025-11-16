package main

import "fmt"

func main() {
	// Deklarasi dan Inisialisasi Variabel
	var nama string = "Didan"
	var umur int = 20
	var tinggi float64 = 175.5
	var isStudent bool = true

	// short variable declaration (recommended)
	alamat := "Comal"
	kota := "Pemalang"

	// Menampilkan Nilai Variabel
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Alamat:", alamat)
	fmt.Println("Kota:", kota)

	fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
	fmt.Printf("Tinggi: %.2f cm, alamat: %s\n", tinggi, alamat)
	fmt.Printf("Mahasiswa: %t\n", isStudent)
}
