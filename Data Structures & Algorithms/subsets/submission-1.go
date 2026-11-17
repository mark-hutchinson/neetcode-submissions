import "slices"

func subsets(nums []int) [][]int {
	results := search(0, nums, nil, nil)
	return results
}

func search(i int, nums []int, currentSet []int, results[][]int) [][]int {
	// base case
	if i == len(nums) {
		clone := slices.Clone(currentSet)
		results = append(results, clone)
		return results
	}

	// Branch 
	// one side has all sets containing num[i]
	newCurrentSet := append(currentSet, nums[i])
	results = search(i+1, nums, newCurrentSet, results)

	// other side has all sets not containing nums[i]
	results = search(i+1, nums, currentSet, results)
	
	return results
}