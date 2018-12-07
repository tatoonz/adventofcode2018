package day4

// --- Part Two ---
// Strategy 2: Of all guards, which guard is most frequently asleep on the same minute?
//
// In the example above, Guard #99 spent minute 45 asleep more than any other guard or minute - three times in total. (In all other cases, any guard spent any minute asleep at most twice.)
//
// What is the ID of the guard you chose multiplied by the minute you chose? (In the above example, the answer would be 99 * 45 = 4455.)

func part2(input []string) int {
	guardMetas := generateGuardMetas(input)

	mostCount := 0

	frequentlySleepMinute := 0
	chooseGuardID := 0

	for guardID, gm := range guardMetas {
		for minute, count := range gm.sleepMinutesCount {
			if count > mostCount {
				mostCount = count

				frequentlySleepMinute = minute
				chooseGuardID = guardID
			}
		}
	}

	return chooseGuardID * frequentlySleepMinute
}
