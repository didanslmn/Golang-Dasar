package main

import "fmt"

func main() {
	fmt.Println("Latihan 1")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d adalah bilangan genap\n", i)
		} else {
			fmt.Printf("%d adalah bilangan ganjil\n", i)
		}
	}

	// segitiga bintang
	n := 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// segitiga bintang kebalik

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// segitiga bintang rata kanan
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// segitiga bintang rata kiri kebalik
	for i := n; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//seigitiga bintang tengah
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// segitiga bintang tengah kebalik
	for i := n; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
