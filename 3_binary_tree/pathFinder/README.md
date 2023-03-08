# Tree Path Finder

Write a function, pathFinder, that takes in the root of a binary tree and a target value. The function should return a slice representing a path to the target value. If the target value is not found in the tree, then return nil.

You may assume that the tree contains unique values.

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

pathFinder(a, "e") // []interface{}{"a", "b", "e"}
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

pathFinder(a, "p") // nil
```

```
Test 3:
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

pathFinder(a, "c") // []interface{}{"a", "c"}
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

pathFinder(a, "h") // []interface{}{"a", "c", "f", "h"}
```

```
Test 5:
x := &Node{Val: "x"}

pathFinder(x, "x") // []interface{}{"x"}
```

```
Test 6:
pathFinder(nil, "x") // nil
```

```
Test 7:
root := &Node{Val: 0}
curr := root

for i := 1; i <= 6000; i++ {
	curr.Right = &Node{Val: i}
	curr = curr.Right
}

//      0
//       \
//        1
//         \
//          2
//           \
//            3
//             .
//              .
//               .
//              5999
//                \
//                6000

pathFinder(root, 3451); // -> []interface{0, 1, 2, 3, ..., 3450, 3451}
```
