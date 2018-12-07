package day5

import (
	"regexp"
	"strings"
)

// --- Part Two ---
// Time to improve the polymer.
//
// One of the unit types is causing problems; it's preventing the polymer from collapsing as much as it should. Your goal is to figure out which unit type is causing the most problems, remove all instances of it (regardless of polarity), fully react the remaining polymer, and measure its length.
//
// For example, again using the polymer dabAcCaCBAcCcaDA from above:
//
// Removing all A/a units produces dbcCCBcCcD. Fully reacting this polymer produces dbCBcD, which has length 6.
// Removing all B/b units produces daAcCaCAcCcaDA. Fully reacting this polymer produces daCAcaDA, which has length 8.
// Removing all C/c units produces dabAaBAaDA. Fully reacting this polymer produces daDA, which has length 4.
// Removing all D/d units produces abAcCaCBAcCcaA. Fully reacting this polymer produces abCBAc, which has length 6.
// In this example, removing all C/c units was best, producing the answer 4.
//
// What is the length of the shortest polymer you can produce by removing all units of exactly one type and fully reacting the result?

func part2(s string) int {
	allStrs := map[string]bool{}

	for _, r := range s {
		ss := string(r)
		ss = strings.ToLower(ss)

		allStrs[ss] = true
	}

	var min int

	for ss := range allStrs {
		re := regexp.MustCompile(`(?i)` + ss)
		normalizedSS := re.ReplaceAllString(s, "")

		units := part1(normalizedSS)

		if min == 0 || units < min {
			min = units
		}
	}

	return min
}
