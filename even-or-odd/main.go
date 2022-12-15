package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter Number:")
	fmt.Scan(&number)

	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println("Even number is", i)
		} else {
			fmt.Println("Odd number is", i)
		}
	}
}
