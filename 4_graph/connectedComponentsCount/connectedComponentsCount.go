package main

type Graph map[int][]int
type Set map[int]bool

func main(){}

func connectedComponentsCount(graph Graph) int {
	count := 0
	visited := Set{}
	for node := range graph {
		if explore(graph, node, visited) {
			count ++
		}
	}
	return count
}

// Recursive Depth First Traversal Util Function for Connected Components Count
func explore(graph Graph, current int, visited Set) bool {
	if visited[current] {
		return false
	}
	visited[current] = true
	for _, neighbor := range graph[current] {
		explore(graph, neighbor, visited)
	}
	return true
}