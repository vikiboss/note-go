package main

import (
	"fmt"
)

func main() {
	arr := []int{}

	for i := 1; i <= 6; i++ {
		arr = append(arr, []int{i}...)
	}

	for i, v := range arr {
		fmt.Printf("Hello Go! value: %d, index: %d\n", v, i)
	}
}
