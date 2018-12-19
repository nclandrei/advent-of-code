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

	scanner := bufio.NewScanner(f)

	freqMap := map[int]struct{}{
		0: struct{}{},
	}

	var resultingFrequency int
	var frequencies []int

	for scanner.Scan() {
		line := scanner.Text()

		frequency, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not parse line into int: %v\n", err)
		}

		resultingFrequency += frequency

		_, ok := freqMap[resultingFrequency]
		if ok {
			log.Printf("first duplicate frequency is: %d\n", resultingFrequency)
			return
		}

		freqMap[resultingFrequency] = struct{}{}

		frequencies = append(frequencies, frequency)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v\n", err)
	}

	for {
		for _, frequency := range frequencies {
			resultingFrequency += frequency

			_, ok := freqMap[resultingFrequency]
			if ok {
				log.Printf("first duplicate frequency is: %d\n", resultingFrequency)
				return
			}

			freqMap[resultingFrequency] = struct{}{}

		}
	}
}
