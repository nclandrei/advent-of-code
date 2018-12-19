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

	var boxIDs []string

	for scanner.Scan() {
		boxID := scanner.Text()
		for _, b := range boxIDs {
			if result, areFabricBoxes := compare(b, boxID); areFabricBoxes {
				log.Printf("fabric box ID: %s\n", result)
				return
			}
		}

		boxIDs = append(boxIDs, boxID)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v\n", err)
	}

	log.Printf("boxes containing fabric have the following runes: %s\n", "")
}

func compare(b1, b2 string) (string, bool) {
	if len(b1) != len(b2) {
		return "", false
	}

	idx := -1

	for i := 0; i < len(b1); i++ {
		if b1[i] == b2[i] {
			continue
		}
		if idx >= 0 {
			return "", false
		}
		idx = i
	}

	if idx < 0 {
		return "", false
	}

	return b1[:idx] + b1[idx+1:], true
}
