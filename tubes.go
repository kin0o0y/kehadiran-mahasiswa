package main

import "fmt"

type mahasiswa struct {
	Nama  string
	Absen [14]string
}

var dataMahasiswa [100]mahasiswa
var jmlmahasiswa int = 20

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
	fmt.Println("===== DATA ABSENSI IF-49-03 =====")

	for i := 0; i < jmlmahasiswa; i++ {

		hadir := hitungHadir(dataMahasiswa[i])
		izin := hitungIzin(dataMahasiswa[i])
		sakit := hitungSakit(dataMahasiswa[i])
		alpa := hitungAlpa(dataMahasiswa[i])

		fmt.Println()
		fmt.Println("Nama :", dataMahasiswa[i].Nama)
		fmt.Println("Hadir :", hadir)
		fmt.Println("Izin :", izin)
		fmt.Println("Sakit :", sakit)
		fmt.Println("Alpa :", alpa)
	}
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
	fmt.Println("0. Keluar")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	return pilih
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
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scan(&nama)

	if pilih == 1 {

		index = sequentialSearch(nama)

	} else if pilih == 2 {

		index = binarySearch(nama)

	}

	if index == -1 {

		fmt.Println("Mahasiswa tidak ditemukan")

	} else {

		fmt.Println()
		fmt.Println("Data ditemukan")
		fmt.Println("Nama :", dataMahasiswa[index].Nama)
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

		if pertemuan >= 1 && pertemuan <= 14 {

			fmt.Println("1. Hadir")
			fmt.Println("2. Izin")
			fmt.Println("3. Sakit")
			fmt.Println("4. Alfa")

			fmt.Print("Pilih status : ")
			fmt.Scan(&status)

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
		} else {

			fmt.Println("Pertemuan tidak valid")

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
	fmt.Println("1. Selection Sort Nama")
	fmt.Println("2. Insertion Sort Nama")
	fmt.Println("3. Selection Sort Hadir Terbanyak")
	fmt.Println("4. Insertion Sort Alfa Terbanyak")
	fmt.Println("5. Selection Sort Izin Terbanyak")
	fmt.Println("6. Insertion Sort Sakit Terbanyak")

	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectSortNama()
		tampilData()
	}

	if pilih == 2 {
		insertSortNama()
		tampilData()
	}

	if pilih == 3 {
		selectSortHadir()
		tampilData()
	}

	if pilih == 4 {
		insertSortAlfa()
		tampilData()
	}

	if pilih == 5 {
		selectSortIzin()
		tampilData()
	}

	if pilih == 6 {
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

	fmt.Print("Masukkan nama mahasiswa : ")
	fmt.Scan(&nama)

	index := sequentialSearch(nama)

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
