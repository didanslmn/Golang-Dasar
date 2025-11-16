package main

import "fmt"

// definisi interface
type Bentuk interface {
	Luas() float64
	Keliling() float64
}

// implementasi struct
type Persegi struct {
	Sisi float64
}

type Lingkaran struct {
	JariJari float64
}

type PersegiPanjang struct {
	Panjang, Lebar float64
}

// implementasi method untuk Persegi (otomatis mengimplementasi interface Bentuk)
func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}
func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// implementasi method untuk Lingkaran (otomatis mengimplementasi interface Bentuk)
func (l Lingkaran) Luas() float64 {
	return 3.14 * l.JariJari * l.JariJari
}
func (l Lingkaran) Keliling() float64 {
	return 2 * 3.14 * l.JariJari
}

// implementasi method untuk PersegiPanjang (otomatis mengimplementasi interface Bentuk)
func (pp PersegiPanjang) Luas() float64 {
	return pp.Panjang * pp.Lebar
}
func (pp PersegiPanjang) Keliling() float64 {
	return 2 * (pp.Panjang + pp.Lebar)
}

// function dengan interface sebagai parameter
func HitungLuas(b Bentuk) {
	fmt.Printf("Luas: %.2f\n", b.Luas())
}

func main() {
	persegi := Persegi{Sisi: 5}
	lingkaran := Lingkaran{JariJari: 7}
	persegiPanjang := PersegiPanjang{Panjang: 4, Lebar: 6}

	// menggunakan function dengan interface
	HitungLuas(persegiPanjang)
	HitungLuas(persegi)
	HitungLuas(lingkaran)

	// emtry interface
	var data interface{} = "Halo, ini adalah empty interface"
	fmt.Println(data)
	data = 12345
	fmt.Println(data)
}
