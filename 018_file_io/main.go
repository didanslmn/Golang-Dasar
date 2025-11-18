package main

import (
	"fmt"
	"os"
)

func main() {
	// buat file
	file, err := os.Create("contoh.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// tulis file
	_, err = file.WriteString("Belajar golang")
	if err != nil {
		panic(err)
	}
	file.Sync()

	// baca file
	data, err := os.ReadFile("contoh.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("isi file:")
	fmt.Println(string(data))

	// hapus file
	err = os.Remove("contoh.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file berhasil dihapus")
}
