# Tree Levels

Write a function, treeLevels, that takes in the root of a binary tree. The function should return a 2-Dimensional slice where each subslice represents a level of the tree.

```
Test 1:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

treeLevels(a)
// [][]interface{}{{"a"}, {"b", "c"}, {"d", "e", "f"}}
```
