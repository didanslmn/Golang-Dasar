package main

import "fmt"

func main() {
	// map
	mahasiswa := map[string]int{
		"andi": 20,
		"budi": 22,
		"caca": 19,
	}

	// menambahkan data ke dalam map
	mahasiswa["dodi"] = 21

	// mengakses data dari map
	fmt.Println("Umur andi:", mahasiswa["andi"])

	// cek keberadaan key dalam map
	umur, ok := mahasiswa["budi"]
	if ok {
		fmt.Println("Umur budi:", umur)
	} else {
		fmt.Println("Data budi tidak ditemukan")
	}

	// menghapus data dari map
	delete(mahasiswa, "caca")

	// iterasi map menggunakan for range
	for nama, umur := range mahasiswa {
		fmt.Printf("nama %s, umur %d\n", nama, umur)
	}
}
