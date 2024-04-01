package main

import (
	"fmt"
	"strconv"
)

// If the number is divisible by 3, write "Foo" otherwise, the number
func Fooer(input int) string {

	isfoo := (input % 3) == 0

	if isfoo {
		return "Foo"
	}

	return strconv.Itoa(input)
}

func Fooer5(input int) string {

	isfoo5 := (input % 5) == 0
	fmt.Println(isfoo5)

	if isfoo5 {
		return "Foo"
	}

	return "Incorrect"
}

func main() {
	output := Fooer5(5)
	fmt.Println(output)
}
