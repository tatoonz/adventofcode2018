// Package day2 part 2
// --- Part Two ---
// Confident that your list of box IDs is complete, you're ready to find the boxes full of prototype fabric.
//
// The boxes will have IDs which differ by exactly one character at the same position in both strings. For example, given the following box IDs:
//
// abcde
// fghij
// klmno
// pqrst
// fguij
// axcye
// wvxyz
// The IDs abcde and axcye are close, but they differ by two characters (the second and fourth). However, the IDs fghij and fguij differ by exactly one character, the third (h and u). Those must be the correct boxes.
//
// What letters are common between the two correct box IDs? (In the example above, this is found by removing the differing character from either ID, producing fgij.)
package day2

import "strings"

func part2(input []string) string {
	for i := 0; i < len(input)-1; i++ {
		a := input[i]

		for j := i + 1; j < len(input); j++ {
			b := input[j]

			missMatch := 0
			missStr := ""
			for ai, as := range a {
				if as != rune(b[ai]) {
					missMatch++
					missStr = string(as)
				}
			}

			if missMatch == 1 {
				return strings.Replace(a, missStr, "", 1)
			}
		}
	}

	return ""
}
