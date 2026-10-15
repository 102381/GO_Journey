package main

import (
	"fmt"
)

func main() {
	age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 50)
	// fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is less than 45")
	}

	names := []string{"mario", "luigi", "goofy", "mickey"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continue at position ", index)
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
		if index > 2 {
			fmt.Printf("breaking at position %v", index)
			break
		}
	}

}
