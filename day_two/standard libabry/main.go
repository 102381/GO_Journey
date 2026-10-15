package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there friends!"

	// // fmt.Println(strings.Contains(greeting, "hello"))
	// // fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Print(strings.Index(greeting, "th"))

	// fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"zero", "one", "two", "three"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "browser"))

}
