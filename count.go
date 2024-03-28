package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	command := os.Args[1]
	filePath := os.Args[2]
	file := file_extraction(filePath)

	if command == "-l" {
		fmt.Println(line_count(file))
	} else if command == "-w" {
		fmt.Println(word_count(file))
	} else if command == "-c" {
		fmt.Println(character_count(file))
	}

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
		str := scanner.Text()
		for i := 0; i < len(str); i++ {
			count++
		}
	}
	return count
}
