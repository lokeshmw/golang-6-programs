package main

import "fmt"

func main() {

	var primNum int = 711
	var primcount int = 0
	for i := 2; i < primNum/2; i++ {
		if primNum%i == 0 {
			primcount++
			break
		}
	}

	if primcount == 0 && primNum != 1 {
		fmt.Println(primNum, " is a Prime Number")
	} else {
		fmt.Println(primNum, " is Not a Prime Number")
	}
}
