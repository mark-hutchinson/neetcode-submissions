func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	for _, p := range prerequisites {
		before := p[1]
		after := p[0]
		adj[before] = append(adj[before], after)
	}

	visited := make(map[int]struct{})
	visiting := make(map[int]struct{})
	for c :=0; c < numCourses; c++{
		if _, ok := visited[c]; ok {
			// already visited
			continue
		}
		cycleDetected := checkCycle(c, adj, visited, visiting)
		if cycleDetected {
			return false
		}

	}
	return true
}

func checkCycle(node int, adj map[int][]int, visited map[int]struct{}, visiting map[int]struct{}) bool {
	neighbours := adj[node]


	if _, ok := visiting[node]; ok {
		return true // cycle detected
	}
	if _, ok := visited[node]; ok {
		return false // visited in a different search
	}
	visited[node] = struct{}{}
	visiting[node] = struct{}{}

	for _, n := range neighbours {
		cycleDetected := checkCycle(n, adj, visited, visiting)
		if cycleDetected {
			return true
		}
	}
	delete(visiting, node)

	
	return false
}