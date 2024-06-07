package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		number := i
		for number != 0 {
			if number%10 == 7 {
				count++
				number /= 10
				fmt.Println(number)
			}
		}
		fmt.Println("lol")
	}
	fmt.Println(count)
}
