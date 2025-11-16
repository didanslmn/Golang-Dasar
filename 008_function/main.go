package main

import (
	"fmt"
	"sync/atomic"
)

// function dengan single return
func tambah(a, b int) int {
	return a + b
}

// function dengan multiple return
func hitung(a, b int) (int, int, int) {
	kali := a * b
	tambah := a + b
	kurang := a - b
	return kali, tambah, kurang
}

// named return values
func hitung2(a, b int) (kali, tambah, kurang int) {
	kali = a * b
	tambah = a + b
	kurang = a - b
	return
}

// anonymous function
func contohAnonymousFunction() {
	counter := int64(0)
	increment := func() {
		atomic.AddInt64(&counter, 1)
	}
	increment()
	increment()
	increment()
	fmt.Println("Counter:", counter)
}

func main() {
	// penggunaan function dengan single return
	hasil := tambah(10, 5)
	fmt.Println("Hasil tambah:", hasil)
	// penggunaan function dengan multiple return
	k, t, kr := hitung(10, 5)
	fmt.Printf("Hasil kali: %d, tambah: %d, kurang: %d\n", k, t, kr)
	// ignore named return values dengan _
	hasilkali, _, _ := hitung2(20, 10)
	fmt.Printf("Hasil kali dari hitung2: %d\n", hasilkali)
	// penggunaan anonymous function
	contohAnonymousFunction()

}
