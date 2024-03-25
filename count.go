package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := os.Args[1]

	fmt.Println(line_count(filePath))
	fmt.Println(word_count(filePath))
}

func line_count(filepath string) int {

	var count int
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
	}
	return count
}

func word_count(filepath string) int {
	count := 0
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count = count + len(strings.Split(scanner.Text(), " "))
	}
	return count
}
