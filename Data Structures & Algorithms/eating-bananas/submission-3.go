func minEatingSpeed(piles []int, h int) int {

	// piles: 1 4 3 2
    // k = 1: 1 4 3 2   = 10
	// k = 2: 1 2 2 1   = 6
	// k = 3: 1 2 1 1   = 5
	// k = 4: 1 1 1 1   = 4
	// worst: k=1
	// best: k=4, max height of piles

	maxVal := 0
	for _, p := range piles{
		if p > maxVal {
			maxVal = p
		}
	}
	minVal :=1

	k := 0
	for minVal < maxVal {
		// test midway between min and max
		k = minVal + (maxVal - minVal)/2
		t := eatingTime(piles, k)
		//fmt.Printf("min=%d, max=%d, k=%d, time=%d\n", minVal, maxVal, k, t)
		if t > h {
			// this K is too slow.
			minVal = k + 1
		} else if t<=h {
			// This value of k works.
			maxVal = k
		}
	}

	// // Brute force. start with k=MaxVal, keep trying lower values of K until eatingTime >h
	// for k:=maxVal; k>0; k-- {
	// 	t := eatingTime(piles, k)
	// 	if t> h {
	// 		return k+1 // the value we just saw was the las tacceptable one
	// 	}
	// }

	
	return minVal
}

func eatingTime(piles []int, k int) int {
	total := 0
	for _, p := range piles {
		t := int(math.Ceil(float64(p)/float64(k)))
		//fmt.Printf("time to eat %d=%d\n", p, t)
		total += t
	}
	return total
}
