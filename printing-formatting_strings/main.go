package main

import "fmt"

func main() {
	age := 45
	name := "Zero"

	//Print
	fmt.Print("Hello,") //doesnt add a new line different from Println
	fmt.Print("World! \n")
	// fmt.Print("my age is", age, " and my name is", name)

	//Printf (formattedstrings) %_ = format specifier
	//Printf also doesnt add new lines
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name) // %q only works on strings
	fmt.Printf("age is o type %T \n", age)                     //gives the type of variable
	fmt.Printf("you scored %0.1f points! \n", 255.55)          //limit float by doing %(num)f

	//sprintf (save formatted strings)

	var str = fmt.Sprintf("my  age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is", str)

}
