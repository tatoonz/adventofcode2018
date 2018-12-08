package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// --- Day 6: Chronal Coordinates ---
// The device on your wrist beeps several times, and once again you feel like you're falling.
//
// "Situation critical," the device announces. "Destination indeterminate. Chronal interference detected. Please specify new target coordinates."
//
// The device then produces a list of coordinates (your puzzle input). Are they places it thinks are safe or dangerous? It recommends you check manual page 729. The Elves did not give you a manual.
//
// If they're dangerous, maybe you can minimize the danger by finding the coordinate that gives the largest distance from the other points.
//
// Using only the Manhattan distance, determine the area around each coordinate by counting the number of integer X,Y locations that are closest to that coordinate (and aren't tied in distance to any other coordinate).
//
// Your goal is to find the size of the largest area that isn't infinite. For example, consider the following list of coordinates:
//
// 1, 1
// 1, 6
// 8, 3
// 3, 4
// 5, 5
// 8, 9
// If we name these coordinates A through F, we can draw them on a grid, putting 0,0 at the top left:
//
// ..........
// .A........
// ..........
// ........C.
// ...D......
// .....E....
// .B........
// ..........
// ..........
// ........F.
// This view is partial - the actual grid extends infinitely in all directions. Using the Manhattan distance, each location's closest coordinate can be determined, shown here in lowercase:
//
// aaaaa.cccc
// aAaaa.cccc
// aaaddecccc
// aadddeccCc
// ..dDdeeccc
// bb.deEeecc
// bBb.eeee..
// bbb.eeefff
// bbb.eeffff
// bbb.ffffFf
// Locations shown as . are equally far from two or more coordinates, and so they don't count as being closest to any.
//
// In this example, the areas of coordinates A, B, C, and F are infinite - while not shown here, their areas extend forever outside the visible grid. However, the areas of coordinates D and E are finite: D is closest to 9 locations, and E is closest to 17 (both including the coordinate's location itself). Therefore, in this example, the size of the largest area is 17.
//
// What is the size of the largest area that isn't infinite?

type coordinate struct {
	x int
	y int
}

type coordinates = []coordinate

func newCoordiate(p string) (coordinate, error) {
	pp := strings.Split(p, ", ")

	c := coordinate{}

	x, err := strconv.Atoi(pp[0])
	if err != nil {
		return c, err
	}

	y, err := strconv.Atoi(pp[1])
	if err != nil {
		return c, err
	}

	c.x = x
	c.y = y

	return c, nil
}

func (c1 coordinate) manhattanDistance(c2 coordinate) int {
	r := math.Abs(float64(c1.x-c2.x)) + math.Abs(float64(c1.y-c2.y))
	return int(r)
}

func part1(ss []string) (int, error) {
	cc := coordinates{}

	maxX := 0
	maxY := 0

	for _, s := range ss {
		c, err := newCoordiate(s)
		if err != nil {
			return 0, err
		}

		maxX = int(math.Max(float64(maxX), float64(c.x)))
		maxY = int(math.Max(float64(maxY), float64(c.y)))

		cc = append(cc, c)
	}

	fmt.Println("max X:", maxX)
	fmt.Println("max Y:", maxY)
	fmt.Println()

	infiniteCoord := map[coordinate]bool{}
	coordClosestCount := map[coordinate]int{}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			// fmt.Printf("C(%v, %v):\n", x, y)

			curC, err := newCoordiate(fmt.Sprintf("%v, %v", x, y))
			if err != nil {
				return 0, err
			}

			closestDist := math.MaxInt64
			closestCoord := coordinate{}
			closestCount := map[int]int{}

			for _, c := range cc {
				d := curC.manhattanDistance(c)

				// fmt.Printf("%v -> %v = %v\n", curC, c, curC.manhattanDistance(c))

				closestCount[d]++

				if d < closestDist {
					closestDist = d
					closestCoord = c
				}
			}

			isOnlyClosestCoord := closestCount[closestDist] == 1
			if isOnlyClosestCoord {
				coordClosestCount[closestCoord]++

				isOnInfiniteEdge := x == 0 || y == 0 || x == maxX || y == maxY
				if isOnInfiniteEdge {
					infiniteCoord[closestCoord] = true
				}
			}
		}
	}

	maxClosestCount := 0
	for coord, count := range coordClosestCount {
		if infiniteCoord[coord] {
			continue
		}

		if count > maxClosestCount {
			maxClosestCount = count
		}
	}

	return maxClosestCount, nil
}
