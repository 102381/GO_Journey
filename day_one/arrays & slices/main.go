package main

import "fmt"

func main() {
	var ages [3]int = [3]int{20, 25, 30} //arrays in go have a fixed lengths
	var agesOne = [3]int{20, 25, 30}     // same as above

	names := [4]string{"zero", "one", "two", "three"}

	fmt.Println(names, len(names))
	fmt.Println(ages, agesOne)

	//slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 45)
	//it doesnt change scores
	//creates new slice
	//overrrides the array

	fmt.Println(scores, len(scores))

	//slice ranges
	//can also be appened
	rangeOne := names[1:3]  //includes first numer exclude second number
	rangeTwo := names[2:]   //begin at 2
	rangeThree := names[:2] //only first 2

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
