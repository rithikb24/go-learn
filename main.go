package main

import "fmt"

func main() {
	c := []int{1, 2, 3, 4, 5}
	d := 8

	fmt.Print(check_existence(c, d))
}

func check_existence(arr []int, d int) bool {
	fmt.Println(arr, d)
	for i := 0; i < len(arr); i++ {
		if d == arr[i] {
			return true
		}
	}
	return false
}
