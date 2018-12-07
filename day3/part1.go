package day3

import (
	"log"
	"regexp"
	"strconv"
)

// --- Day 3: No Matter How You Slice It ---
// The Elves managed to locate the chimney-squeeze prototype fabric for Santa's suit (thanks to someone who helpfully wrote its box IDs on the wall of the warehouse in the middle of the night). Unfortunately, anomalies are still affecting them - nobody can even agree on how to cut the fabric.
//
// The whole piece of fabric they're working on is a very large square - at least 1000 inches on each side.
//
// Each Elf has made a claim about which area of fabric would be ideal for Santa's suit. All claims have an ID and consist of a single rectangle with edges parallel to the edges of the fabric. Each claim's rectangle is defined as follows:
//
// The number of inches between the left edge of the fabric and the left edge of the rectangle.
// The number of inches between the top edge of the fabric and the top edge of the rectangle.
// The width of the rectangle in inches.
// The height of the rectangle in inches.
// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge, 2 inches from the top edge, 5 inches wide, and 4 inches tall. Visually, it claims the square inches of fabric represented by # (and ignores the square inches of fabric represented by .) in the diagram below:
//
// ...........
// ...........
// ...#####...
// ...#####...
// ...#####...
// ...#####...
// ...........
// ...........
// ...........
// The problem is that many of the claims overlap, causing two or more claims to cover part of the same areas. For example, consider the following claims:
//
// #1 @ 1,3: 4x4
// #2 @ 3,1: 4x4
// #3 @ 5,5: 2x2
// Visually, these claim the following areas:
//
// ........
// ...2222.
// ...2222.
// .11XX22.
// .11XX22.
// .111133.
// .111133.
// ........
// The four square inches marked with X are claimed by both 1 and 2. (Claim 3, while adjacent to the others, does not overlap either of them.)
//
// If the Elves all proceed with their own plans, none of them will have enough fabric. How many square inches of fabric are within two or more claims?

type claim struct {
	ID string
	x  int
	y  int
	w  int
	h  int
}

type fabric map[string][]string

func (f fabric) markClaim(c claim) {
	for x := c.x; x < c.x+c.w; x++ {
		for y := c.y; y < c.y+c.h; y++ {
			f[string(x)+string(y)] = append(f[string(x)+string(y)], c.ID)
		}
	}
}

var reClaimPattern = regexp.MustCompile("#(\\d+)\\s@\\s(\\d+),(\\d+):\\s(\\d+)x(\\d+)")

func newClaim(s string) (claim, error) {
	ss := reClaimPattern.FindStringSubmatch(s)

	c := claim{}

	x, err := strconv.Atoi(ss[2])
	if err != nil {
		return c, err
	}

	y, err := strconv.Atoi(ss[3])
	if err != nil {
		return c, err
	}

	w, err := strconv.Atoi(ss[4])
	if err != nil {
		return c, err
	}

	h, err := strconv.Atoi(ss[5])
	if err != nil {
		return c, err
	}

	c.ID = ss[1]
	c.x = x
	c.y = y
	c.w = w
	c.h = h

	return c, nil
}

func part1(input []string) int {
	f := fabric{}

	for _, s := range input {
		c, err := newClaim(s)
		if err != nil {
			log.Fatal(err)
		}

		f.markClaim(c)
	}

	overlapped := 0
	for _, ids := range f {
		if len(ids) > 1 {
			overlapped++
		}
	}

	return overlapped
}
