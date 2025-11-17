package main

import "fmt"

// fungsi kalkulator yang menerima operator dan variadic numbers
func kalkulator(operator string, numbers ...int) int {
	result := 0
	switch operator {
	case "+":
		for _, n := range numbers {
			result += n
		}
	case "-":
		for _, n := range numbers {
			result -= n
		}
	case "*":
		result = 1
		for _, n := range numbers {
			result *= n
		}
	case "/":
		if len(numbers) > 0 {
			result = numbers[0]
			for _, n := range numbers[1:] {
				if n != 0 {
					result /= n
				}
			}
		}
	default:
		fmt.Println("Operator tidak valid")
	}
	return result
}

func main() {
	// contoh penggunaan fungsi kalkulator
	fmt.Println("Penjumlahan:", kalkulator("+", 10, 20, 30))
	fmt.Println("Pengurangan:", kalkulator("-", 100, 30, 20))
	fmt.Println("Perkalian:", kalkulator("*", 2, 3, 4))
	fmt.Println("Pembagian:", kalkulator("/", 100, 2, 5))
}
