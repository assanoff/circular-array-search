package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}
	target := 3

	idx := CircularArraySearch(arr, target)

	if idx == -1 {
		fmt.Printf("array does not contain target %d\n", target)
		return
	}

	fmt.Printf("target is at index %d\n", idx)

}
