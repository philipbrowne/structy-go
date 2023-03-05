# Breadth First Values

Write a function, breadthFirstValues, that takes in the root of a binary tree. The function should return a slice containing all values of the tree in breadth-first order.

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
breadthFirstValues(a)
// []interface{}{"a", "b", "c", "d", "e", "f"}
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
a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
e.Right = h
breadthFirstValues(a)
// []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}
```

```
Test 3:
a := &Node{Val: "a"}
breadthFirstValues(a)
// []interface{}{"a"}
```

```
Test 4:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
x := &Node{Val: "x"}

a.Right = b
b.Left = c
c.Left = x
c.Right = d
d.Right = e

//      a
//       \
//        b
//       /
//      c
//    /  \
//   x    d
//         \
//          e


breadthFirstValues(a)
// []interface{}{"a", "b", "c", "x", "d", "e"}
```

```
Test 5:
breadthFirstValues(nil)
// []interface{}{"a"}
```
