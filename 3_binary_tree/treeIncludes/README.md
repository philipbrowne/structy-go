# Tree Includes

Write a function, treeIncludes, that takes in the root of a binary tree and a target value. The function should return a boolean indicating whether or not the value is contained in the tree.

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

treeIncludes(a, "e") // -> true
```

```
Test 2:
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

treeIncludes(a, "a") // -> true
```

```
Test 3:
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

treeIncludes(a, "n") // -> false
```

```
Test 4:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}
g := &Node{Val: "g"}
h := &Node{Val: "h"}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
//    /       \
//   g         h

treeIncludes(a, "f") // -> true
```

```
Test 5:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}
g := &Node{Val: "g"}
h := &Node{Val: "h"}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
//    /       \
//   g         h

treeIncludes(a, "p") // -> false
```

```
Test 6:
treeIncludes(nil, "b") // -> false
```
