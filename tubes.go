package main

import "fmt"

const NMAX = 100

type RiwayatTidur struct {
	Tanggal     string  // Format: "YYYY-MM-DD"
	JamTidur    string  // Format: "HH:MM"
	JamBangun   string  // Format: "HH:MM"
	DurasiTidur float64 // Dalam jam, hasil perhitungan
}

var dataTidur [NMAX]RiwayatTidur
var jumlahData int = 0

func main() {
	var pilihan int

	for {
		fmt.Println("\n=== APLIKASI PEMANTAUAN KESEHATAN DAN POLA TIDUR ===")
		fmt.Println("1. Tambah Data Tidur")
		fmt.Println("2. Ubah Data Tidur")
		fmt.Println("3. Hapus Data Tidur")
		fmt.Println("4. Cari Data Tidur")
		fmt.Println("5. Urutkan Data Tidur")
		fmt.Println("6. Laporan Pola Tidur")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			ubahData()
		case 3:
			hapusData()
		case 4:
			cariData()
		case 5:
			urutkanData()
		case 6:
			tampilkanLaporan()
		case 0:
			fmt.Println("Terima kasih! Selamat istirahat.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silahkan coba lagi!")
		}
	}
}

func hitungDurasi(jamTidur, jamBangun string) float64 {
	var jam1, menit1, jam2, menit2 int

	fmt.Sscanf(jamTidur, "%d:%d", &jam1, &menit1)
	fmt.Sscanf(jamBangun, "%d:%d", &jam2, &menit2)

	totalMenitTidur := (jam2*60 + menit2) - (jam1*60 + menit1)
	if totalMenitTidur < 0 {
		totalMenitTidur += 24 * 60 // Lewat tengah malam
	}

	return float64(totalMenitTidur) / 60.0
}

func tambahData() {
	if jumlahData >= NMAX {
		fmt.Println("Data sudah penuh.")
		return
	}

	var tanggal, jamTidur, jamBangun string

	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)

	fmt.Print("Masukkan jam tidur (HH:MM): ")
	fmt.Scan(&jamTidur)

	fmt.Print("Masukkan jam bangun (HH:MM): ")
	fmt.Scan(&jamBangun)

	durasi := hitungDurasi(jamTidur, jamBangun)

	dataTidur[jumlahData] = RiwayatTidur{
		Tanggal:     tanggal,
		JamTidur:    jamTidur,
		JamBangun:   jamBangun,
		DurasiTidur: durasi,
	}
	jumlahData++

	fmt.Println("Data berhasil ditambahkan!")

	// Saran/feedback durasi tidur
	if durasi < 7 {
		fmt.Println("⚠ Saran: Durasi tidur Anda kurang dari 7 jam. Usahakan untuk tidur lebih lama demi kesehatan.")
	} else if durasi > 9 {
		fmt.Println("⚠ Saran: Durasi tidur Anda lebih dari 9 jam. Tidur terlalu lama juga kurang baik untuk kesehatan.")
	} else {
		fmt.Println("💬 Feedback: Durasi tidur Anda teratur (7-9 jam). Pertahankan pola tidur ini!")
	}

	// Saran/feedback jam tidur
	var jamTidurJam, jamTidurMenit int
	fmt.Sscanf(jamTidur, "%d:%d", &jamTidurJam, &jamTidurMenit)
	if jamTidurJam >= 23 {
		fmt.Println("⚠ Saran: Anda tidur di atas jam 11 malam. Usahakan untuk tidur lebih awal agar lebih sehat.")
	} else {
		fmt.Println("💬 Feedback: Anda sudah tidur sebelum jam 11 malam. Pertahankan kebiasaan baik ini!")
	}
}

func ubahData() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data untuk diubah.")
		return
	}

	var tanggal string
	fmt.Print("Masukkan tanggal data yang ingin diubah (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)

	// Sequential Search
	idx := -1
	for i := 0; i < jumlahData; i++ {
		if dataTidur[i].Tanggal == tanggal {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Data ditemukan:")
	fmt.Println("Jam Tidur Sebelumnya:", dataTidur[idx].JamTidur)
	fmt.Println("Jam Bangun Sebelumnya:", dataTidur[idx].JamBangun)

	var jamTidurBaru, jamBangunBaru string
	fmt.Print("Masukkan jam tidur baru (HH:MM): ")
	fmt.Scan(&jamTidurBaru)

	fmt.Print("Masukkan jam bangun baru (HH:MM): ")
	fmt.Scan(&jamBangunBaru)

	durasiBaru := hitungDurasi(jamTidurBaru, jamBangunBaru)

	// Update data
	dataTidur[idx].JamTidur = jamTidurBaru
	dataTidur[idx].JamBangun = jamBangunBaru
	dataTidur[idx].DurasiTidur = durasiBaru

	fmt.Println("Data berhasil diperbarui!")
}

func hapusData() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data untuk dihapus.")
		return
	}

	var tanggal string
	fmt.Print("Masukkan tanggal data yang ingin dihapus (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)

	// Sequential Search
	idx := -1
	for i := 0; i < jumlahData; i++ {
		if dataTidur[i].Tanggal == tanggal {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	// Geser elemen ke kiri untuk menghapus
	for i := idx; i < jumlahData-1; i++ {
		dataTidur[i] = dataTidur[i+1]
	}

	jumlahData--

	fmt.Println("Data berhasil dihapus!")
}

func cariData() {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data untuk dicari.")
		return
	}

	var tanggal string
	var metode int

	fmt.Print("Masukkan tanggal yang dicari (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilihan: ")
	fmt.Scan(&metode)

	var idx int = -1

	if metode == 1 {
		// Sequential Search
		for i := 0; i < jumlahData; i++ {
			if dataTidur[i].Tanggal == tanggal {
				idx = i
			}
		}
	} else if metode == 2 {
		// Binary Search
		left := 0
		rigt := jumlahData - 1
		for left <= rigt && idx == -1 {
			mid := (left + rigt) / 2
			if dataTidur[mid].Tanggal == tanggal {
				idx = mid
			} else if dataTidur[mid].Tanggal < tanggal {
				left = mid + 1
			} else {
				rigt = mid - 1
			}
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	if idx != -1 {
		fmt.Println("\nData ditemukan:")
		fmt.Println("Tanggal     :", dataTidur[idx].Tanggal)
		fmt.Println("Jam Tidur   :", dataTidur[idx].JamTidur)
		fmt.Println("Jam Bangun  :", dataTidur[idx].JamBangun)
		fmt.Printf("Durasi Tidur: %.2f jam\n", dataTidur[idx].DurasiTidur)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func urutkanData() {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data untuk diurutkan.")
		return
	}

	var metode, kriteria, urutan int

	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilihan: ")
	fmt.Scan(&metode)

	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Tanggal")
	fmt.Println("2. Durasi Tidur")
	fmt.Print("Pilihan: ")
	fmt.Scan(&kriteria)

	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilihan: ")
	fmt.Scan(&urutan)

	ascending := urutan == 1

	if metode == 1 {
		selectionSort(kriteria, ascending)
	} else if metode == 2 {
		insertionSort(kriteria, ascending)
	} else {
		fmt.Println("Metode tidak valid.")
		return
	}

	fmt.Println("Data berhasil diurutkan!")
}

func selectionSort(kriteria int, ascending bool) {
	for i := 0; i < jumlahData-1; i++ {
		idxTerpilih := i
		for j := i + 1; j < jumlahData; j++ {
			if banding(dataTidur[j], dataTidur[idxTerpilih], kriteria, ascending) {
				idxTerpilih = j
			}
		}
		// Tukar
		dataTidur[i], dataTidur[idxTerpilih] = dataTidur[idxTerpilih], dataTidur[i]
	}
}

func insertionSort(kriteria int, ascending bool) {
	for i := 1; i < jumlahData; i++ {
		temp := dataTidur[i]
		j := i - 1

		for j >= 0 && banding(temp, dataTidur[j], kriteria, ascending) {
			dataTidur[j+1] = dataTidur[j]
			j--
		}
		dataTidur[j+1] = temp
	}
}

func banding(a, b RiwayatTidur, kriteria int, ascending bool) bool {
	if kriteria == 1 { // Tanggal
		if ascending {
			return a.Tanggal < b.Tanggal
		} else {
			return a.Tanggal > b.Tanggal
		}
	} else { // Durasi Tidur
		if ascending {
			return a.DurasiTidur < b.DurasiTidur
		} else {
			return a.DurasiTidur > b.DurasiTidur
		}
	}
}

func tampilkanLaporan() {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data untuk ditampilkan.")
		return
	}

	// Pastikan data diurutkan berdasarkan tanggal ascending
	selectionSort(1, true)

	fmt.Println("===== Laporan Pola Tidur =====")

	// Rekap 7 hari terakhir
	fmt.Println("\n📊 Rekap 7 Hari Terakhir:")
	start := 0
	if jumlahData > 7 {
		start = jumlahData - 7
	}

	total7Hari := 0.0
	jumlahHari := 0

	for i := start; i < jumlahData; i++ {
		fmt.Printf("- %s: %.2f jam\n", dataTidur[i].Tanggal, dataTidur[i].DurasiTidur)
		total7Hari += dataTidur[i].DurasiTidur
		jumlahHari++
	}

	// Rata-rata seluruh minggu
	totalAll := 0.0
	for i := 0; i < jumlahData; i++ {
		totalAll += dataTidur[i].DurasiTidur
	}

	rataRata := totalAll / float64(jumlahData)

	fmt.Printf("\n Rata-rata durasi tidur (semua data): %.2f jam\n", rataRata)

	if jumlahHari > 0 {
		fmt.Printf(" Rata-rata durasi tidur (7 hari terakhir): %.2f jam\n", total7Hari/float64(jumlahHari))
	}
}
