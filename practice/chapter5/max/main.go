package main

import "fmt"

func main() {
	fmt.Println(max(10,2,12,11,5,8))
}

func max(vals...int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}