func minEatingSpeed(piles []int, h int) int {

	// fastest speed - is when K is size of biggest pile
	// slowest speed - k=1
	minK := 1
	maxK := 0
	for _, p := range piles {
		if p > maxK {
			maxK = p
		}
	}

	k := 0
	for minK < maxK {
		k = minK + ((maxK - minK) / 2)
		if k < 1 {
			k = 1
		}
		t := calculateEatingTime(piles, k)
		// fmt.Printf("k=%d [%d, %d], time = %d\n", k, minK, maxK, t)
		if t <= h {
			// k meets the criteria, but is it the minimum k?
			// search smaller k values
			maxK = k
		} else {
			// k is too slow
			minK = k+1
		}
	}
	return minK
}

func calculateEatingTime(piles []int, k int) int {
	t := 0
	for _, p := range piles {
		t += int(math.Ceil(float64(p)/float64(k)))
	}
	return t
}
