package main

// API sederhana yang menerima POST dan menyimpan data ke dala json file
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Product struct {
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

// handler untuk menerima data produk
func productHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// simpan data ke file json
	/*os.O_APPEND → data ditambahkan ke akhir file
	os.O_CREATE → buat file jika belum ada
	os.O_WRONLY → hanya untuk menulis*/
	file, err := os.OpenFile("products.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file) // menulis data + newline ke file
	err = encoder.Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Produk berhasil disimpan")
}

func main() {
	http.HandleFunc("/api/product", productHandler)
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
