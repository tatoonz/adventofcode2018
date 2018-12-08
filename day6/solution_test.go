package day6

import (
	"testing"
)

var exampleInput = deviceInput{
	totalDist: 32,
	coords: []string{
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	},
}

var puzzleInput = deviceInput{
	totalDist: 10000,
	coords: []string{
		"152, 292",
		"163, 90",
		"258, 65",
		"123, 147",
		"342, 42",
		"325, 185",
		"69, 45",
		"249, 336",
		"92, 134",
		"230, 241",
		"74, 262",
		"241, 78",
		"299, 58",
		"231, 146",
		"239, 87",
		"44, 157",
		"156, 340",
		"227, 226",
		"212, 318",
		"194, 135",
		"235, 146",
		"171, 197",
		"160, 59",
		"218, 205",
		"323, 102",
		"290, 356",
		"244, 214",
		"174, 250",
		"70, 331",
		"288, 80",
		"268, 128",
		"359, 98",
		"78, 249",
		"221, 48",
		"321, 228",
		"52, 225",
		"151, 302",
		"183, 150",
		"142, 327",
		"172, 56",
		"72, 321",
		"225, 298",
		"265, 300",
		"86, 288",
		"78, 120",
		"146, 345",
		"268, 181",
		"243, 235",
		"262, 268",
		"40, 60",
	},
}

func TestPart1(t *testing.T) {
	tests := map[int]deviceInput{
		17:   exampleInput,
		3687: puzzleInput,
	}

	for e, i := range tests {
		r, err := solution(i)
		if err != nil {
			t.Error(err)
		}

		if r.largestArea != e {
			t.Errorf("Expected %v, but got: %v\n", e, r.largestArea)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := map[int]deviceInput{
		16:    exampleInput,
		40134: puzzleInput,
	}

	for e, i := range tests {
		r, err := solution(i)
		if err != nil {
			t.Error(err)
		}

		if r.regionSize != e {
			t.Errorf("Expected %v, but got: %v\n", e, r.regionSize)
		}
	}
}
