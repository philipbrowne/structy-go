# Tree Sum

Write a function, treeSum, that takes in the root of a binary tree that contains number values. The function should return the total sum of all values in the tree.

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

treeSum(a)
// 21
```

```
Test 2:
a := &Node{Val: 1}
b := &Node{Val: 6}
c := &Node{Val: 0}
d := &Node{Val: 3}
e := &Node{Val: -6}
f := &Node{Val: 2}
g := &Node{Val: 2}
h := &Node{Val: 2}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//      1
//    /   \
//   6     0
//  / \     \
// 3   -6    2
//    /       \
//   2         2


treeSum(a)
// 10
```

```
Test 3:
treeSum(nil)
// 0
```
