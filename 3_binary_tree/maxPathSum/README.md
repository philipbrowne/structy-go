# Max Root to Leaf Path Sum

Write a function, maxPathSum, that takes in the root of a binary tree that contains number values. The function should return the maximum sum of any root to leaf path within the tree.

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

maxPathSum(a) // -> 18
```

```
Test 2:
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

maxPathSum(a); // -> 59
```

```
Test 3:
a := &Node{Val: -1}
b := &Node{Val: -6}
c := &Node{Val: -5}
d := &Node{Val: -3}
e := &Node{Val: 0}
f := &Node{Val: -13}
g := &Node{Val: -1}
h := &Node{Val: -2}

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
// -3   0    -13
//     /       \
//    -1       -2

maxPathSum(a); // -> -8
```

```
Test 4:
a := &Node{Val: 42}

maxPathSum(a); // -> 42
```
