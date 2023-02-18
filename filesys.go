package main

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fileName string, f func(line string)) {
	if fileName == "" {
		log.Fatal("Error: A filename must be provided.")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}
}
