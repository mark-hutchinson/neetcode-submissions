/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) < 2 {
		return true
	}
	// // O(n^2) - compare each interval
	// for i:=0; i< len(intervals); i++ {
	// 	for j:=i+1; j<len(intervals); j++ {
	// 		if overlaps(intervals[i], intervals[j]) {
	// 			return false
	// 		}
	// 	}
	// }

	// first sort. O(lg n)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].start < intervals[j].start} )

	for i := 0; i < len(intervals) - 1; i++ {
		if intervals[i].end > intervals[i+1].start {
			return false
		}
	}
	return true
}

// func overlaps(one Interval, two Interval) bool {
// 	return one.start < two.end && one.end > two.start
// }
