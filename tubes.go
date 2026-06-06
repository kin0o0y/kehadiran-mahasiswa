package main

import "fmt"

const NMAX int = 100

type mahasiswa struct {
	id, nama                 string
	hadir, izin, alfa, sakit int
}
type status [NMAX]mahasiswa

//function menu utama
func menu() {
	fmt.Println("⌜——————————————————————————————————————————————————————⌝")
	fmt.Println("|SELAMAT DATANG DI SIKEMA (APLIKASI KEHADIRAN MAHASISWA)|")
	fmt.Println("|———————————————————————————————————————————————————————|")
	fmt.Println("| 1. Data Kehadiran Mahasiswa                           |")
	fmt.Println("| 2. Update Kehadiran Mahasiswa                         |")
	fmt.Println("| 3. Cari Data Kehadiran Mahasiswa                      |")
	fmt.Println("| 4. Urutkan Data Kehadiran Mahasiswa                   |")
	fmt.Println("| 5. Statistik Data Kehadiran Mahasiswa                 |")
	fmt.Println("| 6. Keluar                                             |")
	fmt.Println("⌞——————————————————————————————————————————————————————⌟")
	fmt.Print("Pilih menu (1-6): ")
}

//1. function fitur data kehadiran mahasiswa
func isiData(data *status) {
	data[0] = mahasiswa{"001", "Andi", 10, 2, 1, 0}
	data[1] = mahasiswa{"002", "Budi", 12, 0, 0, 1}
	data[2] = mahasiswa{"003", "Citra", 8, 3, 0, 2}
	data[3] = mahasiswa{"004", "Dewi", 15, 0, 0, 0}
	data[4] = mahasiswa{"005", "Eka", 9, 1, 2, 1}
	data[5] = mahasiswa{"006", "Fajar", 11, 1, 1, 1}
	data[6] = mahasiswa{"007", "Gita", 7, 4, 1, 2}
	data[7] = mahasiswa{"008", "Hadi", 13, 0, 1, 1}
	data[8] = mahasiswa{"009", "Intan", 10, 2, 0, 2}
	data[9] = mahasiswa{"010", "Joko", 14, 1, 0, 0}
	data[10] = mahasiswa{"011", "Karin", 8, 2, 3, 1}
	data[11] = mahasiswa{"012", "Lutfi", 12, 0, 2, 1}
	data[12] = mahasiswa{"013", "Maya", 9, 1, 1, 3}
	data[13] = mahasiswa{"014", "Nanda", 15, 0, 0, 0}
	data[14] = mahasiswa{"015", "Oki", 6, 5, 1, 3}
	data[15] = mahasiswa{"016", "Putri", 11, 1, 2, 1}
	data[16] = mahasiswa{"017", "Qori", 13, 0, 1, 1}
	data[17] = mahasiswa{"018", "Raka", 10, 2, 0, 2}
	data[18] = mahasiswa{"019", "Sinta", 8, 3, 1, 3}
	data[19] = mahasiswa{"020", "Tono", 14, 0, 1, 0}
}

//2. function fitur update kehadiran mahasiswa
func updateData(data *status) {
	var nim string
	var pilihan, idx, i int
	idx = -1

	//input nim yang ingin diupdate
	fmt.Print("Masukkan NIM mahasiswa yang ingin diupdate (001 - 020): ")
	fmt.Scan(&nim)

	//cari data mahasiswa berdasarkan nim
	for i = 0; i < 20; i++ {
		if nim == data[i].id {
			idx = i
		}
	}

	//cek nim dan menu update
	if idx == -1 {
		fmt.Println("Data mahasiswa dengan NIM", nim, "tidak ditemukan.")
	} else {
		fmt.Println("Data mahasiswa ditemukan:", data[idx].nama)
		fmt.Println("| Pilih status kehadiran yang ingin diupdate: |")
		fmt.Println("| 1. Hadir                                    |")
		fmt.Println("| 2. Izin                                     |")
		fmt.Println("| 3. Alfa                                     |")
		fmt.Println("| 4. Sakit                                    |")
		fmt.Print("Masukkan pilihan (1-4): ")
		fmt.Scan(&pilihan)
	}
	fmt.Println()

	//update data berdasarkan pilihan
	switch pilihan {
	case 1:
		data[idx].hadir++
	case 2:
		data[idx].izin++
	case 3:
		data[idx].alfa++
	case 4:
		data[idx].sakit++
	default:
		fmt.Println("Pilihan yang diinput tidak valid.")
	}
}

func cariData(data *status) {
	var nim, nama string
	var idx, pilihan, low, high, mid, i int
	idx = -1

	//menu cari data
	fmt.Println("|  Pilih metode pencarian:  |")
	fmt.Println("|  1. Cari berdasarkan Nama |")
	fmt.Println("|  2. Cari berdasarkan NIM  |")
	fmt.Print("Masukkan pilihan (1-2): ")
	fmt.Scan(&pilihan)

	//pilhan pencarian
	switch pilihan {
	case 1:
		fmt.Print("Masukkan Nama Mahasiswa: ")
		fmt.Scan(&nama)
		//cari data mahasiswa berdasarkan nama pakai sequential search
		for i = 0; i < 20; i++ {
			if nama == data[i].nama {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("Data mahasiswa dengan nama", nama, "tidak ditemukan.")
		} else {
			fmt.Println("Data ditemukan:", "ID :", data[idx].id, "| Nama :", data[idx].nama, "| Hadir :", data[idx].hadir, "| Izin :", data[idx].izin, "| Alfa :", data[idx].alfa, "| Sakit :", data[idx].sakit)
		}
	case 2:
		fmt.Print("Masukkan NIM Mahasiswa: ")
		fmt.Scan(&nim)
		//cari data mahasiswa berdasarkan nim
		low = 0
		high = 19
		for low <= high {
			mid = (low + high) / 2
			if data[mid].id == nim {
				idx = mid
				break
			} else if data[mid].id < nim {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if idx == -1 {
			fmt.Println("Data mahasiswa dengan NIM", nim, "tidak ditemukan.")
		} else {
			fmt.Println("Data mahasiswa ditemukan:", "ID :", data[idx].id, "| Nama :", data[idx].nama, "| Hadir :", data[idx].hadir, "| Izin :", data[idx].izin, "| Alfa :", data[idx].alfa, "| Sakit :", data[idx].sakit)
		}
	}
}

func main() {
	var pilih int
	var data status
	menu()
	fmt.Scan(&pilih)
	isiData(&data)
	if pilih == 1 {
		fmt.Println("================== Data Kehadiran Mahasiswa ===================")
		fmt.Println("  ID  |   Nama   |   Hadir   |   Izin   |   Alfa   |   Sakit   ")
		for i := 0; i < 20; i++ {
			fmt.Printf("%-6s| %-9s| %-10d| %-9d| %-9d| %-11d\n", data[i].id, data[i].nama, data[i].hadir, data[i].izin, data[i].alfa, data[i].sakit)
		}
		fmt.Println("===============================================================")
	} else if pilih == 2 {
		updateData(&data)
		fmt.Println("================== Data Kehadiran Mahasiswa ===================")
		fmt.Println("  ID  |   Nama   |   Hadir   |   Izin   |   Alfa   |   Sakit   ")
		for i := 0; i < 20; i++ {
			fmt.Printf("%-6s| %-9s| %-10d| %-9d| %-9d| %-11d\n", data[i].id, data[i].nama, data[i].hadir, data[i].izin, data[i].alfa, data[i].sakit)
		}
		fmt.Println("===============================================================")
	} else if pilih == 3 {
		cariData(&data)
	}
}
