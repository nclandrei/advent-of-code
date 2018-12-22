package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
	Amidst the chaos, you notice that exactly one claim doesn't overlap by even a single square inch of fabric with any other claim.
	If you can somehow draw attention to it, maybe the Elves will be able to make Santa's suit after all!

	For example, in the claims above, only claim 3 is intact after all claims are made.

	What is the ID of the only claim that doesn't overlap?
**/

type claim struct {
	id   int
	left int
	top  int
	wide int
	tall int
}

// #1 @ 1,3: 4x4

func newClaim(line string) (claim, error) {
	parts := strings.Split(line, "@")
	boxID, err := strconv.Atoi(strings.TrimSpace(parts[0])[1:])
	if err != nil {
		return claim{}, err
	}

	subParts := strings.Split(parts[1], ":")
	offsets := strings.Split(strings.TrimSpace(subParts[0]), ",")
	left, err := strconv.Atoi(offsets[0])
	if err != nil {
		return claim{}, err
	}
	top, err := strconv.Atoi(offsets[1])
	if err != nil {
		return claim{}, err
	}

	area := strings.Split(strings.TrimSpace(subParts[1]), "x")
	wide, err := strconv.Atoi(area[0])
	if err != nil {
		return claim{}, err
	}
	tall, err := strconv.Atoi(area[1])
	if err != nil {
		return claim{}, err
	}

	return claim{
		id:   boxID,
		left: left,
		top:  top,
		wide: wide,
		tall: tall,
	}, nil
}

func main() {
	f, err := os.Open("../data/input.txt")
	if err != nil {
		log.Fatalf("could not read from file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		claim, err := newClaim(line)
		if err != nil {
			log.Fatalf("could not convert line into claim: %v\n", err)
		}

		log.Println(claim)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v\n", err)
	}
}

func addClaimToFabricMatrix(claim claim, matrix [][]int) {
	for i := claim.top; i < claim.top+claim.tall; i++ {
		if matrix[i] == nil {
			matrix[i] = make([]int, 1000)
		for j := claim.left; j < claim.left+claim.wide; j++ {
			matrix[i][j] = claim.i
		}
	}
}
