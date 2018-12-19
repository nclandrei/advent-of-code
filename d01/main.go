package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./data/input.txt")
	if err != nil {
		log.Fatalf("could not read from file: %v\n", err)
	}

	scanner := bufio.NewScanner(f)

	var resultingFrequency int
	for scanner.Scan() {
		line := scanner.Text()
		frequency, err := strconv.Atoi(line)
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
