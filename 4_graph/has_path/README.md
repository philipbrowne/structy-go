# Has Path

Write a function, hasPath, that takes in a map representing the adjacency list of a directed acyclic graph and two nodes (src, dst). The function should return a boolean indicating whether or not there exists a directed path between the source and destination nodes.

```
Test 1:
g := Graph{
	"f": []string{"g", "i"},
	"g": []string{"h"},
	"h": []string{},
	"i": []string{"g", "k"},
	"j": []string{"i"},
	"k": []string{},
}

hasPath(g, "f", "k") // true
```

```
Test 2:
g := Graph{
	"f": []string{"g", "i"},
	"g": []string{"h"},
	"h": []string{},
	"i": []string{"g", "k"},
	"j": []string{"i"},
	"k": []string{},
}

hasPath(g, "f", "j") // false
```

```
Test 3:
g := Graph{
	"f": []string{"g", "i"},
	"g": []string{"h"},
	"h": []string{},
	"i": []string{"g", "k"},
	"j": []string{"i"},
	"k": []string{},
}

hasPath(g, "i", "h") // true
```

```
Test 4:
g := Graph{
	"v": []string{"x", "w"},
	"w": []string{},
	"x": []string{},
	"y": []string{"z"},
	"z": []string{},
}

hasPath(g, "v", "w") // true
```

```
Test 5:
g := Graph{
	"v": []string{"x", "w"},
	"w": []string{},
	"x": []string{},
	"y": []string{"z"},
	"z": []string{},
}

hasPath(g, "v", "z") // true
```