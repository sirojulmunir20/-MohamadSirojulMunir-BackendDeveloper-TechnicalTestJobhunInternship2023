# Deskripsi
API ini dibangun untuk mengelola data mahasiswa, beserta jurusan dan hobi yang dimilikinya. Dalam API ini terdapat beberapa endpoint, antara lain:

1. POST /mahasiswa: Endpoint untuk menambahkan data mahasiswa, beserta jurusan dan hobi yang dimilikinya.
2. PUT /mahasiswa/{id}: Endpoint untuk memperbarui data mahasiswa berdasarkan id yang diberikan.
3.GET /mahasiswa: Endpoint untuk mendapatkan daftar seluruh data mahasiswa yang tersimpan.
4. GET /mahasiswa/{id}: Endpoint untuk mendapatkan detail data mahasiswa berdasarkan id yang diberikan.
4. DELETE /mahasiswa/{id}: Endpoint untuk menghapus data mahasiswa berdasarkan id yang diberikan.
Penggunaan


Untuk menggunakan API ini, perlu menjalankan aplikasi Go yang telah dibangun pada server lokal atau server publik yang dapat diakses. Berikut adalah langkah-langkah untuk menjalankan aplikasi:
- Pastikan bahwa komputer atau server yang akan digunakan sudah terinstall Go dan MySQL.
- Clone repository ini ke dalam komputer atau server yang akan digunakan.
- Buatlah database MySQL baru pada server MySQL dengan nama jobhun_test.
- Jalankan skrip SQL yang telah disediakan pada database jobhun_test.
- Konfigurasikan koneksi database pada file main.go.
- Jalankan aplikasi dengan mengetikkan perintah go run main.go pada terminal atau command prompt.