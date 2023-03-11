package main

type Graph map[string][]string

func main () {}

// Depth First Traversal - O(n) Runtime
func hasPathDFI (graph Graph, src, dst string) bool {
	_, ok := graph[src]
	if !ok {
		return false
	}
	_, ok = graph[dst]
	if !ok {
		return false
	}
	stack := []string{src}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == dst {
			return true
		}
		stack = append(stack, graph[current]...)
	}
	return false
}

// Breadth First Traversal - O(n) Runtime
func hasPathBFI (graph Graph, src, dst string) bool {
	_, ok := graph[src]
	if !ok {
		return false
	}
	_, ok = graph[dst]
	if !ok {
		return false
	}
	queue := []string{src}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == dst {
			return true
		}
		queue = append(queue, graph[current]...)
	}
	return false
}

// Depth First Recursive Traversal - O(n) Runtime
func hasPathDFR (graph Graph, src, dst string) bool {
	_, ok := graph[src]
	if !ok {
		return false
	}
	_, ok = graph[dst]
	if !ok {
		return false
	}
	if src == dst {
		return true
	}
	for _, neighbor := range graph[src] {
		if hasPathDFR(graph, neighbor, dst) {
			return true
		}
	}
	return false
}