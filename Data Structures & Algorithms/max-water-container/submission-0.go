func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	max := 0
	for left < right {
		a := area(heights, left, right)
		if a > max {
			max = a
		}
		if heights[right] > heights[left] {
			// keep right where it is
			left++
		} else {
			right--
		}

	}
	return max
}

func area(heights []int, left int, right int) int{
	if right < left {
		fmt.Printf("Right less than left left=%d, right=%d\n", left, right)
		return 0
	}
	width := right - left
	minHeight := heights[left]
	if heights[right] < minHeight {
		minHeight = heights[right]
	}
	return width * minHeight
}
