package main

import "fmt"

func main() {
	// Buat program untuk mengelola daftar belanjaan (slice) dengan fitur tambah, hapus, dan tampilkan.
	daftarBelanja := []string{}

	// Menambahkan item ke daftar belanja
	daftarBelanja = append(daftarBelanja, "Apel")
	daftarBelanja = append(daftarBelanja, "Pisang")
	daftarBelanja = append(daftarBelanja, "Jeruk")

	// hapus item "Pisang" dari daftar belanja
	itemToRemove := "Pisang"
	for i, item := range daftarBelanja {
		if item == itemToRemove {
			daftarBelanja = append(daftarBelanja[:i], daftarBelanja[i+1:]...)
			break
		}
	}

	// menampilkan daftar belanja
	fmt.Println("Daftar Belanja:", daftarBelanja)

	// // Fungsi untuk menambahkan item ke daftar belanja
	// tambahItem := func(item string) {
	// 	daftarBelanja = append(daftarBelanja, item)
	// 	fmt.Println("Item ditambahkan:", item)
	// }
	// // Fungsi untuk menghapus item dari daftar belanja
	// hapusItem := func(item string) {
	// 	for i, v := range daftarBelanja {
	// 		if v == item {
	// 			daftarBelanja = append(daftarBelanja[:i], daftarBelanja[i+1:]...)
	// 			fmt.Println("Item dihapus:", item)
	// 			return
	// 		}
	// 	}
	// 	fmt.Println("Item tidak ditemukan:", item)
	// }
	// // Fungsi untuk menampilkan daftar belanja
	// tampilkanDaftar := func() {
	// 	fmt.Println("Daftar Belanja:")
	// 	for i, item := range daftarBelanja {
	// 		fmt.Printf("%d. %s\n", i+1, item)
	// 	}
	// }
	// // Menambahkan item
	// tambahItem("Apel")
	// tambahItem("Pisang")
	// tambahItem("Jeruk")

	// // Menampilkan daftar belanja
	// tampilkanDaftar()

	// // Menghapus item
	// hapusItem("Pisang")

}
