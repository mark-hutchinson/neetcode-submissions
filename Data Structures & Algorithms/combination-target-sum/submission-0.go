import "slices"

func combinationSum(nums []int, target int) [][]int {


	return search(nums, target, 0, nil, nil)
	
}

func search(nums []int, target int, i int, currentSlice []int, results [][]int) [][]int {
	if target == 0 {
		// currentSlice is an answer. add it to results
		results = append(results, slices.Clone(currentSlice))
		return results
	}
	if target <= 0 {
		// we have gone over
		return results
	}
	if i >= len(nums) {
		return results;
	}



	// consider
	// combinations that include nums[i].  e.g. add nums[i] to current set
	newSlice := append(currentSlice, nums[i])
	results = search(nums, target - nums[i], i, newSlice, results)

	// combinations that do not include nums[i]. consider next number
	results = search(nums, target, i+1, currentSlice, results)


	return results
}
