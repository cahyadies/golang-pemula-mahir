package main

import "fmt"

// Slice potongan dari data Array
func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	// Membuat Slice
	var slice1 = months[4:7]
	// Print slice array
	fmt.Println(slice1)
	// Print panjang slice array
	fmt.Println(len(slice1))
	// Print kapasitas slice array
	fmt.Println(cap(slice1))

	// Iseng ganti data deh
	months[5] = "Adadeh"
	fmt.Println(months)

}
