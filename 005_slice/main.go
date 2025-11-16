package main

import "fmt"

func main() {
	// slice (dinamis)
	buah := []string{"mangga", "apple", "banana"}

	// menambahkan elemen ke dalam slice
	buah = append(buah, "orange", "grape")

	// slice subset
	sliceBUah := buah[1:4] // dari index 1 sampai 3

	fmt.Println("Isi slice buah:", buah)
	fmt.Println("Jumlah buah:", len(buah), "kapasitas:", cap(buah))
	fmt.Println("Jumlah slice buah:", sliceBUah)

	// membuat slice dengan make
	nilai := make([]int, 3, 5) // len 3, cap 5
	nilai = append(nilai, 90, 89, 34)
	fmt.Println("nilai:", nilai)
}
