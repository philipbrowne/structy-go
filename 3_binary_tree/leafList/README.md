# Leaf List

Write a function, leafList, that takes in the root of a binary tree and returns a slice containing the values of all leaf nodes in left-to-right order.

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

leafList(a)
// []interface{}{"d", "e", "f"}
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
f.Right = h

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
//    /       \
//   g         h

leafList(a)
// []interface{}{"d", "g", "h"}
```

```
Test 3:
a := &Node{Val: 5}
b := &Node{Val: 11}
c := &Node{Val: 54}
d := &Node{Val: 20}
e := &Node{Val: 15}
f := &Node{Val: 1}
g := &Node{Val: 3}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
e.Left = f
e.Right = g

//        5
//     /    \
//    11    54
//  /   \
// 20   15
//      / \
//     1  3

leafList(a)
//[]interface{}{20, 1, 3, 54}
```

```
Test 4:
x := &Node{Val: "x"}

leafList(x)
//[]interface{}{"x"}
```

```
Test 5:
leafList(nil)
//[]interface{}{}
```