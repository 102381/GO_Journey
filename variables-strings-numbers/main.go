//learning var, strings, numbers

package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	// strings
	var nameOne string = "Zero"
	var nameTwo = "kingsley"
	var nameThree string

	nameOne = "doc"
	nameThree = "marty"
	nameFour := "Luffy" // short hand version

	//always define string and
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//ints
	var ageOne int = 20
	var ageTWo = 30
	ageThree := 40 // short hand version

	fmt.Println(ageOne, ageTWo, ageThree)

	//bits & memory
	//int 8 ranges from -127 to 128 anything above or below gets an error swap the bit size
	//https://golang.org/ref/spec#Numeric_t...
	// var numOne int8 = 25 //works
	// var NumeTwo int8 = 129 //error

	//uint means u cannot allow a negative number
	var numTwo uint8 = 25
	fmt.Println(numTwo)

	//float
	//specify the bit says is a must
	var scoreOne float32 = 24.54
	var scoreTwo float64 = 12329.3433 //higher precision then 32
	ScoreThree := 1.5                 //is gonna be a float64

	fmt.Println(scoreOne, scoreTwo, ScoreThree)

}
