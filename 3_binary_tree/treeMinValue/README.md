# Tree Minimum Value

Write a function, treeMinValue, that takes in the root of a binary tree that contains number values. The function should return the minimum value within the tree.

You may assume that the input tree is non-empty.

```
Test 1:
a := &Node{Val: 3}
b := &Node{Val: 11}
c := &Node{Val: 4}
d := &Node{Val: 4}
e := &Node{Val: -2}
f := &Node{Val: 1}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//       3
//    /    \
//   11     4
//  / \      \
// 4   -2     1

treeMinValue(a); // -> -2
```

```
Test 2:
a := &Node{Val: 5}
b := &Node{Val: 11}
c := &Node{Val: 3}
d := &Node{Val: 4}
e := &Node{Val: 14}
f := &Node{Val: 12}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//       5
//    /    \
//   11     3
//  / \      \
// 4   14     12

treeMinValue(a); // -> 3
```

```
Test 3:
a := &Node{Val: -1}
b := &Node{Val: -6}
c := &Node{Val: -5}
d := &Node{Val: -3}
e := &Node{Val: -4}
f := &Node{Val: -13}
e := &Node{Val: -2}
f := &Node{Val: -2}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     /       \
//    -2       -2


treeMinValue(a); // -> -13
```

```
Test 4:
a := &Node{Val: 42}

treeMinValue(a); // -> 42
```
