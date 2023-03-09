# All Tree Paths

Write a function, allTreePaths, that takes in the root of a binary tree. The function should return a 2-Dimensional slice where each subslice represents a root-to-leaf path in the tree.

The order within an individual path must start at the root and end at the leaf, but the relative order among paths in the outer slice does not matter.

You may assume that the input tree is non-empty.

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

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f

allTreePaths(a)
// [][]interface{}{{"a", "b", "d"}, {"a", "b", "e"}, {"a", "c", "f"}}
```

```
Test 2:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}
g := &Node{Val: "g"}
h := &Node{Val: "h"}
i := &Node{Val: "i"}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
e.Right = h
f.Left = i

//         a
//      /    \
//     b      c
//   /  \      \
//  d    e      f
//      / \    /
//     g  h   i

allTreePaths(a)
// [][]interface{}{{"a", "b", "d"}, {"a", "b", "e", "g"}, {"a", "b", "e", "h"}, {"a", "c", "f", "i"}}
```

```
Test 3:
q := &Node{Val: "q"}
r := &Node{Val: "r"}
s := &Node{Val: "s"}
tt := &Node{Val: "t"}
u := &Node{Val: "u"}
v := &Node{Val: "v"}

q.Left = r
q.Right = s
r.Right = tt
tt.Left = u
u.Right = v

//      q
//    /   \
//   r     s
//    \
//     t
//    /
//   u
//  /
// v

allTreePaths(q)
// [][]interface{}{{"q", "r", "t", "u", "v"}, {"q", "s"}}
```

```
Test 4:
z := &Node{Val: "z"}

allTreePaths(z)
// [][]interface{}{{"z"}}
```
