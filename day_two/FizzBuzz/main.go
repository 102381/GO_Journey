package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i < n; i++ {

		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "fizzbuzz"
		} else if i%3 == 0 {
			result[i-1] = "buzz"
		} else if i%5 == 0 {
			result[i-1] = "fizz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

func main() {
	n := 100
	result := fizzbuzz(n)
	for _, value := range result {
		fmt.Println(value)
	}
}
