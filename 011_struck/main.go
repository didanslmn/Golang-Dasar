package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NIM     string
	Umur    int
	Jurusan string
}

func main() {
	// instansiasi struct
	mhs1 := Mahasiswa{
		Nama:    "Andi",
		NIM:     "123456",
		Umur:    20,
		Jurusan: "Informatika",
	}

	mhs2 := Mahasiswa{"Didan", "654321", 21, "Sistem Informasi"}

	// menampilkan data mahasiswa
	fmt.Printf("data Mahsasiswa %+v\n", mhs1)
	fmt.Printf("Nama: %s, Jurusan: %s\n", mhs1.Nama, mhs1.Jurusan)

	// mengakses field struct
	fmt.Println("Nama Mahasiswa 1:", mhs1.Nama)
	fmt.Println("NIM Mahasiswa 2:", mhs2.NIM)
}
