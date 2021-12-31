package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Sugeng"
	names[2] = "Cahyadi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(len(names))

}
