package main

import (
	"fmt"
)

func main() {
	fmt.Println("blockchain")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range arr[2:] {
		v = len(arr) - i
		fmt.Println(i, v)
	}

	fmt.Println(arr)
}
