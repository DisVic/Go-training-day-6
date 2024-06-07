package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2 float64
	_, _ = fmt.Scan(&number1, &number2)
	difference := math.Max(number1, number2)
	counter := math.Min(number1, number2)
	var r int = 10
	for r != 0 {
		r = int(difference) - int(difference/counter)*int(counter)
		if r == 0 {
			break
		} else {
			difference = counter
			counter = float64(r)
		}
	}
	fmt.Println(counter)
}
