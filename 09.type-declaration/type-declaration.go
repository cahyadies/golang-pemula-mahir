package main

import "fmt"

func main() {

	//type <nama alias> <type data>
	type noKTP string

	var noKTPEko noKTP = "1234567890"

	fmt.Println(noKTPEko)
	fmt.Println(noKTP("2389"))
}
