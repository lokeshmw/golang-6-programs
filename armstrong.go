package main

import "fmt"

func main() {
	var number = 370
	var tempNumber, remainder int
	var result int = 0
	tempNumber = number
	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10

		if tempNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Printf("%d is an Armstrong number.", number)
	} else {
		fmt.Printf("%d is not an Armstrong number.", number)
	}
}
