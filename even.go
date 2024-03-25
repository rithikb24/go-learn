package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var five_mod []int
	for i := 0; i < len(a); i++ {
		if prime(a[i]) == true {
			fmt.Println("five")
			five_mod = append(five_mod, a[i])
		}
	}
	fmt.Println(five_mod)

}

func prime(a int) bool {
	// divide each number by all numbers trailing upto it and return false if 1
	if a == 1 {
		return false
	}

	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func even(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func odd(a int) bool {
	if a%2 == 0 {
		return false
	}
	return true
}

func five(a int) bool {
	if a%5 == 0 {
		if even(a) == true {
			fmt.Println("here")
			return true
		}
		return false
	}
	return false
}
