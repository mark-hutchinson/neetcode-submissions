func countComponents(n int, edges [][]int) int {

	// map of node -> neighbours
	adj := make(map[int][]int) 
	for _, e := range edges {
		// connections go both ways. (undirected graph)
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visited := make(map[int]bool) // all nodes processed

	count := 0
	for i:= 0; i <n; i++ {
		if visited[i] {
			continue
		}

		// found a new section. count it and explore
		count++
		dfs(i, adj, visited)
	}

	return count
}

func dfs(n int, adj map[int][]int, visited map[int]bool) {
	if visited[n] {
		return
	}

	visited[n] = true
	// visit all adjacent
	for _, a := range adj[n] {
		dfs(a, adj, visited)
	}
}