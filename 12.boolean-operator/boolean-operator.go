package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	var lulusUjian bool = ujian > 75
	var lulusAbsensi bool = absensi > 75
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 75 && absensi >= 75)
}
