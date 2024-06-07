package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var count int
	number := n
	for n != 0 {
		fmt.Scan(&n)
		if n > 0 && number < 0 || n < 0 && number > 0 {
			count++
		}
		number = n
	}
	fmt.Println(count)
}
