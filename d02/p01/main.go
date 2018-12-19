package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../data/input.txt")
	if err != nil {
		log.Fatalf("could not read from file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var boxesWithTwoLetters int
	var boxesWithThreeLetters int

	for scanner.Scan() {
		boxID := scanner.Text()
		hasTwoLetters, hasThreeLetters := doesBoxHaveTwoThreeLetters(boxID)
		if hasTwoLetters {
			boxesWithTwoLetters++
		}
		if hasThreeLetters {
			boxesWithThreeLetters++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v\n", err)
	}

	log.Printf("checksum is: %d\n", boxesWithThreeLetters*boxesWithTwoLetters)
}

func doesBoxHaveTwoThreeLetters(boxID string) (bool, bool) {
	var hasTwoLetters bool
	var hasThreeLetters bool

	charMap := map[rune]int{}
	for _, r := range boxID {
		count, ok := charMap[r]
		if ok {
			count++
			charMap[r] = count
			continue
		}
		charMap[r] = 1
	}

	for _, count := range charMap {
		if hasTwoLetters && hasThreeLetters {
			return hasTwoLetters, hasThreeLetters
		}
		if count == 2 {
			hasTwoLetters = true
		}
		if count == 3 {
			hasThreeLetters = true
		}
	}

	return hasTwoLetters, hasThreeLetters
}
