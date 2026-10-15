package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var inilials []string
	for _, v := range names {
		inilials = append(inilials, v[:1]) //range ussed upon a string appneding to slice
	}

	if len(inilials) > 1 {
		return inilials[0], inilials[1]
	}

	return inilials[0], "_"
}

func main() {
	fn1, ln1 := getInitials("roranoa zoro")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("Tony chopper")
	fmt.Println(fn2, ln2)

	fn3, ln3 := getInitials("Brook")
	fmt.Println(fn3, ln3)
}
