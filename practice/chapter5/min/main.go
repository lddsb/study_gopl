package main

import "fmt"

func main() {
	fmt.Println(min(1, 2, 3, 4, 5, 6, 7, 7))
}

func min(vals ...int) int {
	var min int
	for index, val := range vals {
		if index == 0 {
			min = val
		}
		if min > val {
			min = val
		}
	}
	return min
}
