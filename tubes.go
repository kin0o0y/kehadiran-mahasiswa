package main

import "fmt"

const NMAX int = 100

type mahasiswa struct {
	id, nama                 string
	hadir, izin, alfa, sakit int
}
type status [NMAX]mahasiswa

func menu() {
	fmt.Println("⌜——————————————————————————————————————————————————————⌝")
	fmt.Println("|SELAMAT DATANG DI SIKEMA (APLIKASI KEHADIRAN MAHASISWA)|")
	fmt.Println("————————————————————————————————————————————————————————")
	fmt.Println("| 1. Data Kehadiran Mahasiswa                           |")
	fmt.Println("| 2. Update Kehadiran Mahasiswa                         |")
	fmt.Println("| 3. Cari Data Kehadiran Mahasiswa                      |")
	fmt.Println("| 4. Urutkan Data Kehadiran Mahasiswa                   |")
	fmt.Println("| 5. Statistik Data Kehadiran Mahasiswa                 |")
	fmt.Println("| 6. Keluar                                             |")
	fmt.Println("⌞——————————————————————————————————————————————————————⌟")
	fmt.Print("Pilih menu (1-6): ")
}

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
	}
}
