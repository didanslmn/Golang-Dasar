# Belajar Go Dasar

Selamat datang di repositori untuk belajar dasar-dasar bahasa pemrograman Go! Repositori ini berisi kumpulan kode contoh yang disusun secara sistematis untuk membantu pemula memahami konsep-konsep fundamental di Go.

## Struktur Repositori

Materi pembelajaran dipecah menjadi beberapa bagian yang disimpan dalam direktori bernomor. Setiap direktori berfokus pada satu topik spesifik dan berisi contoh kode yang relevan.

## Topik yang Dibahas

Berikut adalah daftar topik yang dicakup dalam repositori ini, sesuai dengan urutan direktorinya:

- `001_tipe_data_variable`: Pengenalan tipe data dasar (string, int, bool, dll.) dan cara mendeklarasikan variabel.
- `002_control_flow`: Logika kontrol program seperti `if-else`, `switch-case`, dan perulangan (`for`).
- `003_latihan_1`: Latihan pertama untuk menguji pemahaman konsep-konsep awal.
- `004_array`: Penggunaan tipe data koleksi `array` yang memiliki ukuran tetap.
- `005_slice`: Pengenalan `slice`, tipe data koleksi yang lebih fleksibel dan dinamis.
- `006_map`: Penggunaan `map` untuk menyimpan data dalam format key-value.
- `007_latihan_2`: Latihan kedua yang mencakup array, slice, dan map.
- `008_function`: Cara mendefinisikan dan memanggil fungsi.
- `009_variadic_func`: Pengenalan _variadic function_, yaitu fungsi dengan jumlah argumen yang tidak terbatas.
- `010_latihan_3`: Latihan ketiga seputar penggunaan fungsi.
- `011_struct`: Cara membuat tipe data kustom Anda sendiri menggunakan `struct`.
- `012_method`: Menambahkan fungsi (method) ke dalam sebuah `struct`.
- `013_latihan_4`: Latihan keempat yang berfokus pada struct dan method.
- `014_interface`: Konsep `interface` untuk mendefinisikan "kontrak" atau perilaku.
- `015_error_handling`: Pola penanganan error yang umum di Go.
- `016_package_module`: Penggunaan package dan module untuk mengorganisir kode.
- `017_http_json`: Membuat server HTTP dan bekerja dengan JSON.
- `018_file_io`: Membaca dan menulis file.
- `019_latihan_5`: Latihan kelima untuk menguji pemahaman konsep-konsep lanjutan.
- `API_CRUD`: Contoh aplikasi CRUD (Create, Read, Update, Delete) dengan API.

## Cara Menggunakan

1.  **Clone Repositori (jika belum)**:
    ```bash
    git clone <URL_REPOSITORI_ANDA>
    cd go-dasar
    ```

2. **Navigasi ke Direktori Topik**:
    Pilih topik yang ingin Anda pelajari dan masuk ke direktorinya.
    ```bash
    cd 001_tipe_data_variable
    ```

3. **Jalankan Kode**:
    Jika di dalam direktori tersebut terdapat file Go (misalnya, `main.go`), Anda bisa menjalankannya untuk melihat output dan memahami cara kerjanya.
    ```bash
    go run .
    ```
    Atau jika ada beberapa file, Anda bisa menjalankannya secara spesifik:
    ```bash
    go run nama_file.go
    ```

Selamat belajar!
