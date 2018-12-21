package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
	The Elves managed to locate the chimney-squeeze prototype fabric for Santa's suit
	(thanks to someone who helpfully wrote its box IDs on the wall of the warehouse in the middle of the night).
	Unfortunately, anomalies are still affecting them - nobody can even agree on how to cut the fabric.

	The whole piece of fabric they're working on is a very large square - at least 1000 inches on each side.

	Each Elf has made a claim about which area of fabric would be ideal for Santa's suit.
	All claims have an ID and consist of a single rectangle with edges parallel to the edges of the fabric.

	Each claim's rectangle is defined as follows:
	The number of inches between the left edge of the fabric and the left edge of the rectangle.
	The number of inches between the top edge of the fabric and the top edge of the rectangle.
	The width of the rectangle in inches. The height of the rectangle in inches.
	A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge,
	2 inches from the top edge, 5 inches wide, and 4 inches tall.
	Visually, it claims the square inches of fabric represented by # (and ignores the square inches of fabric represented by .)
	in the diagram below:

	...........
	...........
	...#####...
	...#####...
	...#####...
	...#####...
	...........
	...........
	...........
	The problem is that many of the claims overlap, causing two or more claims to cover part of the same areas.
	For example, consider the following claims:

	#1 @ 1,3: 4x4
	#2 @ 3,1: 4x4
	#3 @ 5,5: 2x2
	Visually, these claim the following areas:

	........
	...2222.
	...2222.
	.11XX22.
	.11XX22.
	.111133.
	.111133.
	........
	The four square inches marked with X are claimed by both 1 and 2. (Claim 3, while adjacent to the others,
	does not overlap either of them.)

	If the Elves all proceed with their own plans, none of them will have enough fabric.
	How many square inches of fabric are within two or more claims?
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
