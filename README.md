# NapTime: Aplikasi Pemantauan Kesehatan dan Pola Tidur Sederhana

## Deskripsi
Aplikasi ini membantu pengguna mencatat dan menganalisis pola tidur serta aktivitas kesehatan harian.  
Data utama yang digunakan adalah **riwayat tidur**, **jam tidur dan bangun**, serta **kualitas tidur pengguna**.  
Pengguna aplikasi adalah individu yang ingin memantau kesehatan tidur mereka.

## Spesifikasi

**a.** Pengguna dapat **menambahkan**, **mengubah**, dan **menghapus** riwayat tidur yang mencakup **jam tidur** dan **jam bangun**.  
**b.** Sistem akan secara otomatis **menghitung durasi tidur** dan memberikan **saran pola tidur yang sehat**.  
**c.** Pengguna dapat **mencari riwayat tidur berdasarkan tanggal** menggunakan **Sequential dan Binary Search**.  
**d.** Pengguna dapat **mengurutkan riwayat tidur** berdasarkan **durasi** atau **tanggal** menggunakan **Selection dan Insertion Sort**.  
**e.** Sistem menampilkan laporan yang mencakup, misalnya:
- Rekapitulasi **durasi tidur dalam 7 hari terakhir**.
- **Rata-rata durasi tidur per minggu**.

## Fitur

- Tambah, ubah, dan hapus data tidur harian
- Hitung durasi tidur otomatis
- Laporan 7 hari terakhir dan rata-rata tidur
- Pencarian data (sequential dan binary search)
- Pengurutan data (selection dan insertion sort)

## Catatan Teknis
- Maksimal data tidur yang disimpan: 100 entri (konstanta NMAX)
- Sistem melakukan validasi jam tidur dan jam bangun.
- Sistem memberikan saran **jika tidur lewat dari jam 23.00** atau **durasi kurang dari 8 jam**

## Cara Menjalankan

Pastikan Anda telah menginstal Go di sistem Anda.  
Lalu jalankan aplikasi dengan perintah berikut di terminal:

```bash
go run naptime_app.go
```
