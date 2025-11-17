package main

import "fmt"

// fungsi variadic adalah fungsi yang dapat menerima jumlah argumen yang tidak terbatas
func tambahangka(angka ...int) int {
	sum := 0
	for _, n := range angka {
		sum += n
	}
	return sum
}

// fungsi variadic dengan argumen tambahan
func sapa(nama string, pesan ...string) {
	fmt.Printf("Halo, %s!\n", nama)
	for _, pesan := range pesan {
		fmt.Println(pesan)
	}
}

// function as parameter
func proses(data []int, fn func(int) int) []int {
	result := make([]int, len(data))
	for i, v := range data {
		result[i] = fn(v)
	}
	return result
}

func main() {
	// contoh penggunaan fungsi variadic
	total := tambahangka(10, 20, 30, 40, 50)
	fmt.Println("Total:", total)

	// contoh penggunaan fungsi variadic dengan argumen tambahan
	sapa("Andi", "Selamat pagi!", "Semoga harimu menyenangkan.")
	sapa("Budi")

	// contoh penggunaan function as parameter
	numbers := []int{1, 2, 3, 4, 5}
	squared := proses(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared:", squared)

	// first-class function
	kaliDua := func(n int) int {
		return n * 2
	}
	angka := []int{1, 2, 3, 4, 5}
	hasil := proses(angka, kaliDua)
	fmt.Println("Hasil kali dua:", hasil)
}
