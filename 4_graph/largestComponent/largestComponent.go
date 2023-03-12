package main

type Graph map[int][]int
type Set map[int]bool

func main(){}

func largestComponent(g Graph) int {
	largest := 0
	visited := Set{}
	for node := range g {
		size := explore(g, node, visited)
		if size > largest {
			largest = size
		} 
	}
	return largest
}

func explore(g Graph, current int, visited Set) int {
	if visited[current] {
		return 0
	}
	visited[current] = true
	size := 1
	for _, neighbor := range g[current] {
		size += explore(g, neighbor, visited)
	}
	return size
}

