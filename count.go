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

	// fmt.Println(line_count(filePath))
	// character_count(filePath)
	file := file_extraction(filePath)
	fmt.Println(line_count(file))
	defer file.Close()
}

func file_extraction(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return file
}

func line_count(file *os.File) int {

	var count int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
	}
	return count
}

func word_count(file *os.File) int {
	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count = count + len(strings.Split(scanner.Text(), " "))
	}
	return count
}

func character_count(file *os.File) int {
	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(str[0]); i++ {
			count++
			fmt.Println(count)
		}
	}
	return count
}
