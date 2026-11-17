func hammingWeight(n int) int {
	count := 0
	for n!= 0 {
		//fmt.Printf("%b\n", n)
		count = count + (n & 1)
		n = n >> 1
	}
	return count
}
