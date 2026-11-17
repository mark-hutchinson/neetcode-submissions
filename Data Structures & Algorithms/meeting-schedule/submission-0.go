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
	// n squared - compare each interval
	for i:=0; i< len(intervals); i++ {
		for j:=i+1; j<len(intervals); j++ {
			if overlaps(intervals[i], intervals[j]) {
				return false
			}
		}
	}
	return true
}

func overlaps(one Interval, two Interval) bool {
	return one.start < two.end && one.end > two.start
}
