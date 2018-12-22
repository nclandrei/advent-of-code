package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	fabricMatrix := make([][]int, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		claim, err := newClaim(line)
		if err != nil {
			log.Fatalf("could not convert line into claim: %v\n", err)
		}

		addClaimToFabricMatrix(claim, fabricMatrix)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning: %v\n", err)
	}

	var overlappingSquaresCount int
	for _, row := range fabricMatrix {
		for _, column := range row {
			if column > 1 {
				overlappingSquaresCount++
			}
		}
	}

	log.Printf("number of overlapping squares is: %d\n", overlappingSquaresCount)
}

func addClaimToFabricMatrix(claim claim, matrix [][]int) {
	for i := claim.top; i < claim.top+claim.tall; i++ {
		if matrix[i] == nil {
			matrix[i] = make([]int, 1000)
		}
		for j := claim.left; j < claim.left+claim.wide; j++ {
			matrix[i][j]++
		}
	}
}
