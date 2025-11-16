package main

import "fmt"

type Segitiga struct {
	Alas, Tinggi float64
}

// Method dengan value Reciver (tidak mengubah nilai asli)
func (s Segitiga) luas() float64 {
	return 0.5 * s.Alas * s.Tinggi
}

// Method dengan pointer Reciver (mengubah nilai asli)
func (s *Segitiga) Scale(factor float64) {
	s.Alas *= factor
	s.Tinggi *= factor
}

func main() {
	s := Segitiga{Alas: 4, Tinggi: 5}
	fmt.Printf("Luas segitiga: %.2f\n", s.luas())

	s.Scale(2)
	fmt.Printf("setelah scale - alas: %.2f, tinggi: %.2f\n", s.Alas, s.Tinggi)
	fmt.Printf("Luas segitiga setelah diskalakan: %.2f\n", s.luas())
}
