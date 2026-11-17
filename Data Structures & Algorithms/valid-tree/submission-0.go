func validTree(n int, edges [][]int) bool {
    if n == 0 {
		return true
	}
	adjMap := make(map[int][]int)
	for _, edge := range edges {
		adjMap[edge[0]] = append(adjMap[edge[0]], edge[1])
		adjMap[edge[1]] = append(adjMap[edge[1]], edge[0])
	}

	prev := -1
	current := 0
	visited := make(map[int]struct{})

	cycleDetected := visitNode(current, prev, adjMap, visited)
	if cycleDetected {
		return false
	}

	visitCount := len(visited)
	return visitCount == n
}

func visitNode(current int, prev int, adjMap map[int][]int, visited map[int]struct{}) bool {
	_, ok := visited[current]
	if ok {
		return true // cycle detected
	}

	// visit this node
	visited[current] = struct{}{}

	// visit all children without going backwards
	for _, n := range adjMap[current] {
		if n == prev {
			continue
		}
		cycleDetected := visitNode(n, current, adjMap, visited)
		if cycleDetected {
			return true
		}
	}
	return false
}