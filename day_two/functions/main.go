package main

import (
	"fmt"
	"math"
)

func hellotheir(n string) {
	fmt.Printf("Hello there %v \n", n)
}
func grievous(n string) {
	fmt.Printf("%v my enemies", n)
}

func cyclenames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// hellotheir("general grievous")
	// grievous("general kenobi")
	cyclenames([]string{"luffy", "zoro", "blackbeard"}, hellotheir)
	cyclenames([]string{"luffy", "zoro", "blackbeard"}, grievous)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2 )
}
