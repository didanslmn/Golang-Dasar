package main

import "fmt"

type Buku struct {
	Judul     string
	Pengarang string
	Harga     float64
}

func (b *Buku) Diskon() float64 {
	return b.Harga - (b.Harga * 10 / 100)
}

// func (b *Buku) Diskon(persen float64) float64 {
// 	return b.Harga - (b.Harga * persen / 100)
// }

func main() {
	buku1 := Buku{
		Judul:     "Belajar Go",
		Pengarang: "Andi",
		Harga:     100000,
	}
	fmt.Printf("Harga buku sebelum diskon: %.2f\n", buku1.Harga)
	hargaSetelahDiskon := buku1.Diskon()
	fmt.Printf("Harga buku setelah diskon: %.2f\n", hargaSetelahDiskon)
	// harga_baru := buku1.Diskon(10)
	// fmt.Printf("Harga buku setelah diskon: %.2f\n", harga_baru)
}
