# Depth First Values

Write a function, depthFirstValues, that takes in the root of a binary tree. The function should return a slice containing all values of the tree in depth-first order.

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

depthFirstValues(a)
// []interface{}{"a", "b", "d", "e", "c", "f"}
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

depthFirstValues(a)
// []interface{}{"a", "b", "d", "e", "g", "c", "f"}
```

```
Test 3:
a := &Node{Val: "a"}
depthFirstValues(a)
// []interface{}{"a"}
```

```
Test 4:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}

a.Right = b
b.Left = c
c.Right = d
d.Right = e

//      a
//       \
//        b
//       /
//      c
//       \
//        d
//         \
//          e

depthFirstValues(a)
// []interface{}{"a", "b", "c", "d", "e"}
```

```
Test 5:
depthFirstValues(nil)
// []interface{}{}
```
