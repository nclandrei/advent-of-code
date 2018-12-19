package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../data/input.txt")
	if err != nil {
		log.Fatalf("could not read from file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var resultingFrequency int

	for scanner.Scan() {
		frequency, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("could not parse number: %v\n", err)
		}

		resultingFrequency += frequency
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v\n", err)
	}

	log.Printf("resulting frequency is: %d\n", resultingFrequency)
}
