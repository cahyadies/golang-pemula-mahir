package main

import "fmt"

func main() {
	var name = "Eko"
	var e byte = name[2]
	var eString string = string(e)

	fmt.Println(eString)

}
