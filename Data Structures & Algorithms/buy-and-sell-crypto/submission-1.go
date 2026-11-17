func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	left := 0
	right := 1
	
	maxProfit := 0

	for right < n{
		if prices[right] >= prices[left] {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
			right = right + 1
		} else {
			// Found a new minimum buy value
			// advance left pointer to this new minimum
			left = right
			right = left + 1
		}

	}
	return maxProfit
}
