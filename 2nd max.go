package main

import "fmt"

func main() {
	var max1 int = 0
	var max2 int = 0
	var arr [5]int

	arr[0] = 6
	arr[1] = 4
	arr[2] = 7
	arr[3] = 3
	arr[4] = 2

	max1 = arr[0]
	for i := 1; i <= 4; i++ {
		if max1 < arr[i] {
			max2 = max1
			max1 = arr[i]
		} else if max2 < arr[i] {
			max2 = arr[i]
		}
	}
	fmt.Println("Second largest element is: ", max2)
}
