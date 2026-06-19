package main

import "fmt"

type mahasiswa struct {
	NIM      int
	Nama     string
	Angkatan int
	Aktif    bool
	Absen    [14]string
}

var dataMahasiswa [100]mahasiswa
var jmlmahasiswa int = 20
var dataKeluar [100]mahasiswa
var jmlKeluar int = 0

func inisialisasiData() {
	nama := [20]string{

		"Ahmad_Fauzan",
		"Bima_Pratama",
		"Citra_Lestari",
		"Dinda_Maharani",
		"Eko_Saputra",
		"Farhan_Akbar",
		"Gita_Ramadhani",
		"Hadi_Nugraha",
		"Intan_Permata",
		"Neila_Ezri",
		"Karin_Amelia",
		"Lutfi_Hidayat",
		"Maya_Salsabila",
		"Nanda_Putra",
		"Oki_Setiawan",
		"Putri_Azzahra",
		"Raka_Firmansyah",
		"Sinta_Rahma",
		"Kinasih_Sekar",
		"Vina_Kartika",
	}
	for i := 0; i < jmlmahasiswa; i++ {
		dataMahasiswa[i].Nama = nama[i]
		dataMahasiswa[i].NIM = 130100 + i
		dataMahasiswa[i].Angkatan = 2025
		dataMahasiswa[i].Aktif = true

		for j := 0; j < 14; j++ {
			dataMahasiswa[i].Absen[j] = "H"
		}
	}

}

func hitungHadir(m mahasiswa) int {
	var hadir int
	hadir = 0
	for i := 0; i < 14; i++ {
		if m.Absen[i] == "H" {
			hadir++
		}
	}
	return hadir
}

func hitungIzin(m mahasiswa) int {
	var izin int
	izin = 0

	for i := 0; i < 14; i++ {

		if m.Absen[i] == "I" {
			izin++
		}

	}

	return izin
}

func hitungSakit(m mahasiswa) int {

	sakit := 0

	for i := 0; i < 14; i++ {

		if m.Absen[i] == "S" {
			sakit++
		}

	}

	return sakit
}

func hitungAlpa(m mahasiswa) int {
	alpa := 0
	for i := 0; i < 14; i++ {

		if m.Absen[i] == "A" {
			alpa++
		}

	}

	return alpa
}

func tampilData() {
	fmt.Println()
	fmt.Println("================================================================================")
	fmt.Printf("%-20s %-8s %-10s %-6s %-6s %-6s %-6s %-6s\n",
		"Nama", "NIM", "Angkatan", "Aktif", "Hadir", "Izin", "Sakit", "Alfa")
	fmt.Println("================================================================================")

	for i := 0; i < jmlmahasiswa; i++ {
		fmt.Printf("%-20s %-8d %-10d %-6t %-6d %-6d %-6d %-6d\n",
			dataMahasiswa[i].Nama,
			dataMahasiswa[i].NIM,
			dataMahasiswa[i].Angkatan,
			dataMahasiswa[i].Aktif,
			hitungHadir(dataMahasiswa[i]),
			hitungIzin(dataMahasiswa[i]),
			hitungSakit(dataMahasiswa[i]),
			hitungAlpa(dataMahasiswa[i]))
	}
	fmt.Println("================================================================================")
}

func menuUtama() int {

	var pilih int

	fmt.Println()
	fmt.Println("===== SISTEM ABSENSI IF-49-03 =====")
	fmt.Println("1. Lihat Data")
	fmt.Println("2. Edit Absensi")
	fmt.Println("3. Searching")
	fmt.Println("4. Sorting")
	fmt.Println("5. Statistik")
	fmt.Println("6. Tambah Mahasiswa")
	fmt.Println("7. Hapus Mahasiswa")
	fmt.Println("8. Mahasiswa Keluar")
	fmt.Println("0. Keluar")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih < 0 || pilih > 8 {

		fmt.Println("Menu tidak tersedia")
		return -1

	}

	return pilih
}

func sudahAda(nama string) bool {

	for i := 0; i < jmlmahasiswa; i++ {

		if dataMahasiswa[i].Nama == nama {
			return true
		}

	}

	return false
}

func namaValid(nama string) bool {

	for i := 0; i < len(nama); i++ {

		if nama[i] >= '0' && nama[i] <= '9' {
			return false
		}

	}

	return true
}

func main() {

	inisialisasiData()

	var pilih int = -1

	for pilih != 0 {

		pilih = menuUtama()

		if pilih == 1 {
			tampilData()
		}
		if pilih == 3 {
			menuSearching()
		}
		if pilih == 2 {
			inputAbsensi()
		}
		if pilih == 4 {
			menuSorting()
		}
		if pilih == 5 {
			menuStatistik()
		}
		if pilih == 6 {
			tambahMahasiswa()
		}

		if pilih == 7 {
			hapusMahasiswa()
		}

		if pilih == 8 {
			tampilMahasiswaKeluar()
		}

	}

	fmt.Println("Program selesai")
}

func sequentialSearch(nama string) int {
	index := -1

	for i := 0; i < jmlmahasiswa; i++ {

		if dataMahasiswa[i].Nama == nama {
			index = i
		}

	}

	return index
}

func menuSearching() {

	var pilih int
	var nama string
	var index int

	fmt.Println()
	fmt.Println("===== MENU SEARCHING =====")
	fmt.Println("1. Cari Mahasiswa")
	fmt.Println("2. Cari Mahasiswa Cepat")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scan(&nama)

	if pilih == 1 {
		index = sequentialSearch(nama)

	} else if pilih == 2 {
		insertSortNama()
		index = binarySearch(nama)

	}

	if index == -1 {

		fmt.Println("Mahasiswa tidak ditemukan")

	} else {

		fmt.Println()
		fmt.Println("Data ditemukan")
		fmt.Println("Nama :", dataMahasiswa[index].Nama)
		fmt.Println("NIM :", dataMahasiswa[index].NIM)
		fmt.Println("Angkatan :", dataMahasiswa[index].Angkatan)
		fmt.Println("Aktif :", dataMahasiswa[index].Aktif)
		fmt.Println("Hadir :", hitungHadir(dataMahasiswa[index]))
		fmt.Println("Izin :", hitungIzin(dataMahasiswa[index]))
		fmt.Println("Sakit :", hitungSakit(dataMahasiswa[index]))
		fmt.Println("Alpa :", hitungAlpa(dataMahasiswa[index]))

	}

}

func inputAbsensi() {

	var nama string
	var index int
	var pertemuan int
	var status int

	fmt.Println()
	fmt.Println("===== INPUT ABSENSI =====")

	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scan(&nama)

	index = sequentialSearch(nama)

	if index == -1 {

		fmt.Println("Mahasiswa tidak ditemukan")

	} else {

		fmt.Print("Masukkan pertemuan (1-14) : ")
		fmt.Scan(&pertemuan)

		if pertemuan < 1 || pertemuan > 14 {
			fmt.Println("Pertemuan harus 1 - 14")
			return
		}

		fmt.Println("1. Hadir")
		fmt.Println("2. Izin")
		fmt.Println("3. Sakit")
		fmt.Println("4. Alfa")

		fmt.Print("Pilih status : ")
		fmt.Scan(&status)

		if status < 1 || status > 4 {

			fmt.Println("Status tidak valid")
			return

		}

		if status == 1 {
			dataMahasiswa[index].Absen[pertemuan-1] = "H"
		} else if status == 2 {
			dataMahasiswa[index].Absen[pertemuan-1] = "I"
		} else if status == 3 {
			dataMahasiswa[index].Absen[pertemuan-1] = "S"
		} else if status == 4 {
			dataMahasiswa[index].Absen[pertemuan-1] = "A"
		} else {
			fmt.Println("Status tidak valid")
		}

		fmt.Println("Absensi berhasil diperbarui")
		fmt.Println()
		fmt.Println("Absensi", dataMahasiswa[index].Nama)

		for i := 0; i < 14; i++ {
			fmt.Print("P", i+1, ":", dataMahasiswa[index].Absen[i], " ")
		}

		fmt.Println()

	}
}

func tambahMahasiswa() {

	if jmlmahasiswa >= 100 {

		fmt.Println("Kapasitas penuh")
		return

	}

	var m mahasiswa

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&m.NIM)

	if m.NIM <= 0 {

		fmt.Println("NIM tidak valid")
		return

	}

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&m.Nama)

	if !namaValid(m.Nama) {

		fmt.Println("Nama tidak boleh mengandung angka")
		return

	}

	if sudahAda(m.Nama) {

		fmt.Println("Mahasiswa sudah ada")
		return

	}

	fmt.Print("Masukkan Angkatan : ")
	fmt.Scan(&m.Angkatan)

	if m.Angkatan < 2020 || m.Angkatan > 2030 {

		fmt.Println("Angkatan tidak valid")
		return

	}

	m.Aktif = true

	for i := 0; i < 14; i++ {
		m.Absen[i] = "H"
	}

	dataMahasiswa[jmlmahasiswa] = m
	jmlmahasiswa++

	fmt.Println("Mahasiswa berhasil ditambahkan")
}

func hapusMahasiswa() {

	var nama string

	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scan(&nama)

	index := sequentialSearch(nama)

	if index == -1 {

		fmt.Println("Mahasiswa tidak ditemukan")

	} else {

		dataKeluar[jmlKeluar] = dataMahasiswa[index]
		jmlKeluar++

		for i := index; i < jmlmahasiswa-1; i++ {

			dataMahasiswa[i] = dataMahasiswa[i+1]

		}

		jmlmahasiswa--
		dataMahasiswa[jmlmahasiswa] = mahasiswa{}

		fmt.Println("Mahasiswa berhasil dihapus")
		fmt.Println("Data dipindahkan ke daftar mahasiswa keluar")

	}

}

func tampilMahasiswaKeluar() {

	fmt.Println()
	fmt.Println("===== MAHASISWA KELUAR =====")

	if jmlKeluar == 0 {

		fmt.Println("Belum ada mahasiswa keluar")

	} else {

		for i := 0; i < jmlKeluar; i++ {

			fmt.Println()
			fmt.Println("NIM :", dataKeluar[i].NIM)
			fmt.Println("Nama :", dataKeluar[i].Nama)

		}

	}

}

func selectSortNama() {
	var min int
	var temp mahasiswa

	for i := 0; i < jmlmahasiswa-1; i++ {

		min = i

		for j := i + 1; j < jmlmahasiswa; j++ {

			if dataMahasiswa[j].Nama < dataMahasiswa[min].Nama {
				min = j
			}

		}

		temp = dataMahasiswa[i]
		dataMahasiswa[i] = dataMahasiswa[min]
		dataMahasiswa[min] = temp

	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama (A-Z)")
}

func insertSortNama() {
	var temp mahasiswa
	var j int

	for i := 1; i < jmlmahasiswa; i++ {
		temp = dataMahasiswa[i]
		j = i - 1
		for j >= 0 && dataMahasiswa[j].Nama > temp.Nama {
			dataMahasiswa[j+1] = dataMahasiswa[j]
			j--
		}
		dataMahasiswa[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama (A-Z)")
}

func menuSorting() {

	var pilih int

	fmt.Println()
	fmt.Println("===== MENU SORTING =====")
	fmt.Println("1. Urutkan Berdasarkan Nama (A-Z)")
	fmt.Println("2. Urutkan Berdasarkan Hadir Terbanyak")
	fmt.Println("3. Urutkan Berdasarkan Alfa Terbanyak")
	fmt.Println("4. Urutkan Berdasarkan Izin Terbanyak")
	fmt.Println("5. Urutkan Berdasarkan Sakit Terbanyak")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectSortNama()
		tampilData()
	}

	if pilih == 2 {
		selectSortHadir()
		tampilData()
	}

	if pilih == 3 {
		insertSortAlfa()
		tampilData()
	}

	if pilih == 4 {
		selectSortIzin()
		tampilData()
	}

	if pilih == 5 {
		insertSortSakit()
		tampilData()
	}
}

func selectSortHadir() {
	var max int
	var temp mahasiswa

	for i := 0; i < jmlmahasiswa-1; i++ {

		max = i

		for j := i + 1; j < jmlmahasiswa; j++ {

			if hitungHadir(dataMahasiswa[j]) > hitungHadir(dataMahasiswa[max]) {
				max = j
			}

		}

		temp = dataMahasiswa[i]
		dataMahasiswa[i] = dataMahasiswa[max]
		dataMahasiswa[max] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan hadir terbanyak")
}

func selectSortIzin() {
	var max int
	var temp mahasiswa

	for i := 0; i < jmlmahasiswa-1; i++ {

		max = i

		for j := i + 1; j < jmlmahasiswa; j++ {

			if hitungIzin(dataMahasiswa[j]) > hitungIzin(dataMahasiswa[max]) {
				max = j
			}

		}

		temp = dataMahasiswa[i]
		dataMahasiswa[i] = dataMahasiswa[max]
		dataMahasiswa[max] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan izin terbanyak")
}

func insertSortSakit() {
	var temp mahasiswa
	var j int

	for i := 1; i < jmlmahasiswa; i++ {

		temp = dataMahasiswa[i]

		j = i - 1

		for j >= 0 && hitungSakit(dataMahasiswa[j]) < hitungSakit(temp) {

			dataMahasiswa[j+1] = dataMahasiswa[j]

			j--
		}

		dataMahasiswa[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan sakit terbanyak")
}

func insertSortAlfa() {
	var temp mahasiswa
	var j int

	for i := 1; i < jmlmahasiswa; i++ {

		temp = dataMahasiswa[i]

		j = i - 1

		for j >= 0 && hitungAlpa(dataMahasiswa[j]) < hitungAlpa(temp) {

			dataMahasiswa[j+1] = dataMahasiswa[j]

			j--
		}

		dataMahasiswa[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan alfa terbanyak")
}

func binarySearch(nama string) int {

	var kiri int = 0
	var kanan int = jmlmahasiswa - 1
	var tengah int

	for kiri <= kanan {

		tengah = (kiri + kanan) / 2

		if dataMahasiswa[tengah].Nama == nama {

			return tengah

		} else if dataMahasiswa[tengah].Nama < nama {

			kiri = tengah + 1

		} else {

			kanan = tengah - 1

		}

	}

	return -1
}

func hadirTerbanyak() {

	var max int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungHadir(dataMahasiswa[i]) > hitungHadir(dataMahasiswa[max]) {
			max = i
		}

	}

	fmt.Println()
	fmt.Println("===== HADIR TERBANYAK =====")
	fmt.Println("Nama :", dataMahasiswa[max].Nama)
	fmt.Println("Jumlah Hadir :", hitungHadir(dataMahasiswa[max]))
}

func hadirTersedikit() {

	var min int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungHadir(dataMahasiswa[i]) < hitungHadir(dataMahasiswa[min]) {
			min = i
		}

	}

	fmt.Println()
	fmt.Println("===== HADIR TERSEDIKIT =====")
	fmt.Println("Nama :", dataMahasiswa[min].Nama)
	fmt.Println("Jumlah Hadir :", hitungHadir(dataMahasiswa[min]))
}

func izinTerbanyak() {

	var max int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungIzin(dataMahasiswa[i]) > hitungIzin(dataMahasiswa[max]) {
			max = i
		}

	}

	fmt.Println()
	fmt.Println("===== IZIN TERBANYAK =====")
	fmt.Println("Nama :", dataMahasiswa[max].Nama)
	fmt.Println("Jumlah Izin :", hitungIzin(dataMahasiswa[max]))
}

func izinTersedikit() {

	var min int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungIzin(dataMahasiswa[i]) < hitungIzin(dataMahasiswa[min]) {
			min = i
		}

	}

	fmt.Println()
	fmt.Println("===== IZIN TERSEDIKIT =====")
	fmt.Println("Nama :", dataMahasiswa[min].Nama)
	fmt.Println("Jumlah Izin :", hitungIzin(dataMahasiswa[min]))
}

func sakitTerbanyak() {

	var max int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungSakit(dataMahasiswa[i]) > hitungSakit(dataMahasiswa[max]) {
			max = i
		}

	}

	fmt.Println()
	fmt.Println("===== SAKIT TERBANYAK =====")
	fmt.Println("Nama :", dataMahasiswa[max].Nama)
	fmt.Println("Jumlah Sakit :", hitungSakit(dataMahasiswa[max]))
}

func sakitTersedikit() {

	var min int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungSakit(dataMahasiswa[i]) < hitungSakit(dataMahasiswa[min]) {
			min = i
		}

	}

	fmt.Println()
	fmt.Println("===== SAKIT TERSEDIKIT =====")
	fmt.Println("Nama :", dataMahasiswa[min].Nama)
	fmt.Println("Jumlah Sakit :", hitungSakit(dataMahasiswa[min]))
}

func alpaTerbanyak() {

	var max int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungAlpa(dataMahasiswa[i]) > hitungAlpa(dataMahasiswa[max]) {
			max = i
		}

	}

	fmt.Println()
	fmt.Println("===== ALFA TERBANYAK =====")
	fmt.Println("Nama :", dataMahasiswa[max].Nama)
	fmt.Println("Jumlah Alfa :", hitungAlpa(dataMahasiswa[max]))
}

func alpaTersedikit() {

	var min int = 0

	for i := 1; i < jmlmahasiswa; i++ {

		if hitungAlpa(dataMahasiswa[i]) < hitungAlpa(dataMahasiswa[min]) {
			min = i
		}

	}

	fmt.Println()
	fmt.Println("===== ALFA TERSEDIKIT =====")
	fmt.Println("Nama :", dataMahasiswa[min].Nama)
	fmt.Println("Jumlah Alfa :", hitungAlpa(dataMahasiswa[min]))
}

func statistikMahasiswa() {

	var nama string
	var index int

	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scan(&nama)

	index = sequentialSearch(nama)

	if index == -1 {

		fmt.Println("Mahasiswa tidak ditemukan")

	} else {

		fmt.Println()
		fmt.Println("===== STATISTIK MAHASISWA =====")
		fmt.Println("Nama :", dataMahasiswa[index].Nama)

		fmt.Println("Hadir :", hitungHadir(dataMahasiswa[index]))
		fmt.Println("Izin :", hitungIzin(dataMahasiswa[index]))
		fmt.Println("Sakit :", hitungSakit(dataMahasiswa[index]))
		fmt.Println("Alfa :", hitungAlpa(dataMahasiswa[index]))
	}
}

func menuStatistik() {
	var pilih int

	fmt.Println()
	fmt.Println("===== MENU STATISTIK =====")
	fmt.Println("1. Hadir Terbanyak")
	fmt.Println("2. Hadir Tersedikit")
	fmt.Println("3. Izin Terbanyak")
	fmt.Println("4. Izin Tersedikit")
	fmt.Println("5. Sakit Terbanyak")
	fmt.Println("6. Sakit Tersedikit")
	fmt.Println("7. Alpa Terbanyak")
	fmt.Println("8. Alpa Tersedikit")
	fmt.Println("9. Statistik Mahasiswa")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)
	
	if pilih == 1 {
		hadirTerbanyak()
	}
	if pilih == 2 {
		hadirTersedikit()
	}
	if pilih == 3 {
		izinTerbanyak()
	}
	if pilih == 4 {
		izinTersedikit()
	}
	if pilih == 5 {
		sakitTerbanyak()
	}
	if pilih == 6 {
		sakitTersedikit()
	}
	if pilih == 7 {
		alpaTerbanyak()
	}
	if pilih == 8 {
		alpaTersedikit()
	}
	if pilih == 9 {
		statistikMahasiswa()
	}
}
