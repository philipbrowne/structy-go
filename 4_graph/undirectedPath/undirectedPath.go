package main

type Graph map[string][]string
type Set map[string]bool

func main () {}

func buildGraph(edges [][]string) Graph {
	graph := Graph{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
 	}
	return graph
}

func undirectedPathDFI(edges [][]string, nodeA, nodeB string) bool {
	graph := buildGraph(edges)
	stack := []string{nodeA}
	visited := Set{}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == nodeB {
			return true
		}
		visited[current] = true
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
	return false
}

func undirectedPathBFI(edges [][]string, nodeA, nodeB string) bool {
	graph := buildGraph(edges)
	queue := []string{nodeA}
	visited := Set{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == nodeB {
			return true
		}
		visited[current] = true
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}

func undirectedPathDFR(edges [][]string, nodeA, nodeB string) bool {
	graph := buildGraph(edges)
	return hasPath(graph, nodeA, nodeB, Set{})
}

// Utility Recursive Function for Undirected Path
func hasPath(graph Graph, src, dst string, visited Set) bool {
	if src == dst {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true
	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst, visited) {
			return true
		}
	}
	return false
}