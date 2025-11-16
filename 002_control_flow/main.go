package main

import "fmt"

func main() {
	// Contoh penggunaan if-else
	nilai := 80
	if nilai >= 80 {
		fmt.Println("A")
	} else if nilai >= 70 {
		fmt.Println("B")
	} else if nilai >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// Contoh penggunaan switch-case
	hari := "Rabu"
	switch hari {
	case "Senin":
		fmt.Println("Hari Senin")
	case "Selasa":
		fmt.Println("Hari Selasa")
	case "Rabu":
		fmt.Println("Hari Rabu")
	default:
		fmt.Println("Hari lainnya")
	}

	// Contoh penggunaan for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// Contoh penggunaan break dan continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println("Perulangan ke-", i)
	}

	// Contoh penggunaan nested loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// for range
	bahasa := []string{"go", "python", "java", "javascript"}
	for index, value := range bahasa {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	// tanpa index
	for _, value := range bahasa {
		fmt.Printf("value: %s\n", value)
	}

	// map dengan for range
	person := map[string]string{
		"name":    "Didan",
		"address": "Comal",
		"city":    "Pemalang",
	}
	for key, value := range person {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	// tanpa value
	for key := range person {
		fmt.Printf("key: %s\n", key)
	}
	// tanpa key
	for _, value := range person {
		fmt.Printf("value: %s\n", value)
	}

}
