// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {

	result := make([][]Pair, 0)

	for i :=0; i < len(pairs); i++ {
		
		// everything before i is already sorted
		for j := i-1; j>=0 && pairs[j+1].Key < pairs[j].Key; j-- {
			tmp := pairs[j + 1]
			pairs[j+1] = pairs[j]
			pairs[j] = tmp
		}
		cloned := make([]Pair, len(pairs))
		copy(cloned, pairs)
		result = append(result, cloned)
	}

	return result
}
