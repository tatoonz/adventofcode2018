package day3

import (
	"log"
)

// --- Part Two ---
// Amidst the chaos, you notice that exactly one claim doesn't overlap by even a single square inch of fabric with any other claim. If you can somehow draw attention to it, maybe the Elves will be able to make Santa's suit after all!
//
// For example, in the claims above, only claim 3 is intact after all claims are made.
//
// What is the ID of the only claim that doesn't overlap?
func part2(input []string) string {
	f := fabric{}

	for _, s := range input {
		c, err := newClaim(s)
		if err != nil {
			log.Fatal(err)
		}

		f.markClaim(c)
	}

	overlap := map[string]bool{}
	for _, ids := range f {
		o := false
		if len(ids) > 1 {
			o = true
		}

		for _, id := range ids {
			if overlap[id] == false {
				overlap[id] = o
			}
		}
	}

	for ID, o := range overlap {
		if o == false {
			return ID
		}
	}

	return ""
}
