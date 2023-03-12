# Largest Component

Write a function, largestComponent, that takes in the adjacency list of an undirected graph. The function should return the size of the largest connected component in the graph.

```
Test 1:
g := Graph{
	0: []int{8,1,5},
	1: []int{0},
	5: []int{0,8},
	8: []int{0,5},
	2: []int{3,4},
	3: []int{2,4},
	4: []int{3,2},
}
largestComponent(g) // 4
```

```
Test 2:
g := Graph{
	1: []int{2},
	2: []int{1,8},
	6: []int{7},
	9: []int{8},
	7: []int{6,8},
	8: []int{9,7,2},
}
largestComponent(g) // 6
```

```
Test 3:
g := Graph{
	3: []int{},
	4: []int{6},
	6: []int{4,5,7,8},
	8: []int{6},
	7: []int{6},
	5: []int{6},
	1: []int{2},
	2: []int{1},
}
largestComponent(g) // 5
```

```
Test 4:
g := Graph{}
largestComponent(g) // 0
```

```
Test 5:
g := Graph{
	0: []int{4,7},
	1: []int{},
	2: []int{},
	3: []int{6},
	4: []int{0},
	6: []int{3},
	7: []int{0},
	8: []int{},
}
largestComponent(g) // 3
```