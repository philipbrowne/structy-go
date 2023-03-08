# Tree Value Count

Write a function, tree_value_count, that takes in the root of a binary tree and a target value. The function should return the number of times that the target occurs in the tree.

```
Test 1:
a := &Node{Val: 12}
b := &Node{Val: 6}
c := &Node{Val: 6}
d := &Node{Val: 4}
e := &Node{Val: 6}
f := &Node{Val: 12}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//      12
//    /   \
//   6     6
//  / \     \
// 4   6     12

treeValueCount(a,  6) // 3
```

```
Test 2:
a := &Node{Val: 12}
b := &Node{Val: 6}
c := &Node{Val: 6}
d := &Node{Val: 4}
e := &Node{Val: 6}
f := &Node{Val: 12}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//      12
//    /   \
//   6     6
//  / \     \
// 4   6     12

treeValueCount(a,  12) // 2
```

```
Test 3:
a := &Node{Val: 7}
b := &Node{Val: 5}
c := &Node{Val: 1}
d := &Node{Val: 1}
e := &Node{Val: 8}
f := &Node{Val: 7}
g := &Node{Val: 1}
h := &Node{Val: 1}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//      7
//    /   \
//   5     1
//  / \     \
// 1   8     7
//    /       \
//   1         1

treeValueCount(a,  1) // 4
```

```
Test 4:
a := &Node{Val: 7}
b := &Node{Val: 5}
c := &Node{Val: 1}
d := &Node{Val: 1}
e := &Node{Val: 8}
f := &Node{Val: 7}
g := &Node{Val: 1}
h := &Node{Val: 1}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//      7
//    /   \
//   5     1
//  / \     \
// 1   8     7
//    /       \
//   1         1

treeValueCount(a,  9) // 0
```

```
Test 5:
treeValueCount(nil,  42) // 0
```