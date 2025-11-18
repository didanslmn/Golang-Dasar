package main

import (
	"fmt"
	"module/helper"
)

func main() {
	hasil := helper.Tambah(10, 5)
	fmt.Println("Hasil Tambah:", hasil)

	hasilKurang := helper.Kurang(10, 5)
	fmt.Println("Hasil Kurang:", hasilKurang)

	fmt.Println("haloo:", helper.ToUpper("Didan"))
}
