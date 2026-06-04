package main

import "fmt"

const NMAX int = 100

type mahasiswa struct {
	id, nama                 string
	hadir, izin, alfa, sakit int
}
type status [NMAX]mahasiswa

func main() {
	var n int
	var mhs status	
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mhs[i].id, &mhs[i].nama, &mhs[i].hadir, &mhs[i].izin, &mhs[i].alfa, &mhs[i].sakit)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%-8s | %-11s | %d\n", mhs[i].id, mhs[i].nama, mhs[i].alfa)
	}
}
