# How High

Write a function, howHigh, that takes in the root of a binary tree. The function should return a number representing the height of the tree.

The height of a binary tree is defined as the maximal number of edges from the root node to any leaf node.

If the tree is empty, return -1.

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

howHigh(a) // 2
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

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
//    /
//   g

howHigh(a) // 3
```

```
Test 3:
a := &Node{Val: "a"}
c := &Node{Val: "c"}

a.Right = c

//      a
//       \
//        c

howHigh(a) // 1
```

```
Test 4:
a := &Node{Val: "a"}
//      a

howHigh(a) // 0
```

```
Test 5:
howHigh(nil) // -1
```