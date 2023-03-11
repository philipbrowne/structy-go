# Undirected Path

Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.

```
Test 1:
edges := [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}
	
undirectedPath(edges, "j", "m") // true
```

```
Test 2:
edges := [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}
	
undirectedPath(edges, "m", "j") // true
```

```
Test 3:
edges := [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}
	
undirectedPath(edges, "l", "j") // true
```

```
Test 4:
edges := [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}
	
undirectedPath(edges, "k", "o") // false
```

```
Test 5:
edges := [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}
	
undirectedPath(edges, "i", "o") // false
```

```
Test 6:
edges := [][]string{
	{"b", "a"},
	{"c", "a"},
	{"b", "c"},
	{"q", "r"},
	{"q", "s"},
	{"q", "u"},
	{"q", "t"},
}
	
undirectedPath(edges, "a", "b") // true
```

```
Test 7:
edges := [][]string{
	{"b", "a"},
	{"c", "a"},
	{"b", "c"},
	{"q", "r"},
	{"q", "s"},
	{"q", "u"},
	{"q", "t"},
}
	
undirectedPath(edges, "a", "c") // true
```

```
Test 8:
edges := [][]string{
	{"b", "a"},
	{"c", "a"},
	{"b", "c"},
	{"q", "r"},
	{"q", "s"},
	{"q", "u"},
	{"q", "t"},
}
	
undirectedPath(edges, "r", "t") // true
```

```
Test 9:
edges := [][]string{
	{"b", "a"},
	{"c", "a"},
	{"b", "c"},
	{"q", "r"},
	{"q", "s"},
	{"q", "u"},
	{"q", "t"},
}
	
undirectedPath(edges, "r", "b") // false
```

```
Test 10:
edges := [][]string{
	{"s", "r"},
	{"t", "q"},
	{"q", "r"},
}
undirectedPath(edges, "r", "t") // true
```